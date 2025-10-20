package otel

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	Tracer = otel.Tracer(schemaName)
	Logger = otelslog.NewLogger(schemaName)
)

const schemaName = "https://github.com/grafana/docker-otel-lgtm"

func SetupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	if !viper.GetBool("USE_OLTP_GRAFANA") {
		return
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(prop)

	var traceExporter *otlptrace.Exporter

	if viper.GetBool("OTLP_TRACE_USE_HTTPS") {
		traceExporter, err = otlptrace.New(ctx, otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(viper.GetString("OTLP_TRACE_HTTP")),
		))
	} else {
		traceExporter, err = otlptrace.New(ctx, otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(viper.GetString("OTLP_TRACE_HTTP")),
			otlptracehttp.WithInsecure(),
		))
	}

	if err != nil {
		return nil, err
	}

	tracerProvider := trace.NewTracerProvider(trace.WithBatcher(traceExporter))
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	var metricExporter *otlpmetrichttp.Exporter

	if viper.GetBool("OTLP_METRIC_USE_HTTPS") {
		metricExporter, err = otlpmetrichttp.New(ctx, otlpmetrichttp.WithEndpoint(viper.GetString("OTLP_METRIC_HTTP")))
	} else {
		metricExporter, err = otlpmetrichttp.New(ctx, otlpmetrichttp.WithEndpoint(viper.GetString("OTLP_METRIC_HTTP")), otlpmetrichttp.WithInsecure())
	}

	if err != nil {
		Logger.ErrorContext(ctx, "failed to create metric exporter:", slog.Any("error", err))
		fmt.Println(err)
		return nil, err
	}

	res, resErr := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(viper.GetString("OTEL_SERVICE_NAME")),
			semconv.ServiceInstanceIDKey.String(viper.GetString("OTEL_INSTANCE_ID")),
		),
	)
	if resErr != nil {
		return nil, fmt.Errorf("failed to create resource: %w", resErr)
	}

	metricExportInterval := viper.GetDuration("OTEL_METRIC_EXPORT_INTERVAL_IN_S")
	if metricExportInterval == 0 {
		metricExportInterval = 10 * time.Second
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(metricExporter, metric.WithInterval(metricExportInterval))), metric.WithResource(res))
	if err != nil {
		Logger.ErrorContext(ctx, "failed to create meter provider:", slog.Any("error", err))
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	var logExporter *otlploghttp.Exporter

	if viper.GetBool("OTLP_LOG_USE_HTTPS") {
		logExporter, err = otlploghttp.New(ctx, otlploghttp.WithEndpoint(viper.GetString("OTLP_LOG_HTTP")))
	} else {
		logExporter, err = otlploghttp.New(ctx, otlploghttp.WithEndpoint(viper.GetString("OTLP_LOG_HTTP")), otlploghttp.WithInsecure())
	}

	if err != nil {
		Logger.ErrorContext(ctx, "failed to create log exporter:", slog.Any("error", err))
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(
			log.NewBatchProcessor(
				logExporter,
			),
		),
		log.WithResource(res),
	)

	if err != nil {
		Logger.ErrorContext(ctx, "failed to create log provider:", slog.Any("error", err))
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	err = runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))
	if err != nil {
		Logger.ErrorContext(ctx, "otel runtime instrumentation failed:", slog.Any("error", err))
	}

	return
}
