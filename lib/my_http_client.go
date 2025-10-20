package lib

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
)

type MyHttpClient struct {
	Client         *http.Client
	Builder        *MyHttpClientBuilder
	ResponseString string
}

func (b *MyHttpClient) Execute() {
	logger.Info(b.Builder.Ctx, "URL", b.Builder.Url)
	logger.Info(b.Builder.Ctx, "METHOD", b.Builder.Method)
	logger.Info(b.Builder.Ctx, "HEADERS", b.Builder.Headers)
	logger.Info(b.Builder.Ctx, "PARAMS", b.Builder.Params)

	if b.Builder.VerifySsl {
		err := b.loadCertSsl()
		if err != nil {
			logger.Error(b.Builder.Ctx, "error load cert", map[string]interface{}{
				"error": err,
				"tags":  []string{"net http", "load cert"},
			})
			return
		}
	}

	resp, err := b.Client.Get(b.Builder.Url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	b.ResponseString = string(body)
	fmt.Println("Response:", b.ResponseString)
}

func (b *MyHttpClient) loadCertSsl() error {
	clientCert, err := tls.LoadX509KeyPair(b.Builder.VerifySslClientCrtPath, b.Builder.VerifySslClientKeyPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to load client certificate and key: %v", err))
	}

	// Load CA certificate (to verify the server's certificate)
	caCert, err := os.ReadFile(b.Builder.VerifySslCaCrtPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to read CA certificate: %v", err))
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert}, // Client certificate
		RootCAs:      caCertPool,                    // CA certificate to verify the server
	}

	b.Client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return nil
}

type HttpMethod string

var (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
	OPTION HttpMethod = "OPTION"
)

type MyHttpClientBuilder struct {
	Ctx                    context.Context
	Url                    string
	Headers                map[string]string
	Params                 map[string]interface{}
	Method                 HttpMethod
	Timeout                int
	SaveLog                bool
	VerifySsl              bool
	VerifySslClientCrtPath string
	VerifySslClientKeyPath string
	VerifySslCaCrtPath     string
}

func NewMyHttpClientBuilder() *MyHttpClientBuilder {
	return &MyHttpClientBuilder{}
}

func (b *MyHttpClientBuilder) SetCtx(ctx context.Context) *MyHttpClientBuilder {
	b.Ctx = ctx
	return b
}

func (b *MyHttpClientBuilder) SetUrl(url string) *MyHttpClientBuilder {
	b.Url = url
	return b
}

func (b *MyHttpClientBuilder) SetHeaders(headers map[string]string) *MyHttpClientBuilder {
	b.Headers = headers
	return b
}

func (b *MyHttpClientBuilder) SetParams(params map[string]interface{}) *MyHttpClientBuilder {
	b.Params = params
	return b
}

func (b *MyHttpClientBuilder) SetMethod(method HttpMethod) *MyHttpClientBuilder {
	b.Method = method
	return b
}

func (b *MyHttpClientBuilder) SetTimeout(timeout int) *MyHttpClientBuilder {
	b.Timeout = timeout
	return b
}

func (b *MyHttpClientBuilder) SetSaveLog(saveLog bool) *MyHttpClientBuilder {
	b.SaveLog = saveLog
	return b
}

func (b *MyHttpClientBuilder) SetVerifySsl(verifySsl bool) *MyHttpClientBuilder {
	b.VerifySsl = verifySsl
	return b
}

func (b *MyHttpClientBuilder) SetVerifySslClientCrtPath(verifySslClientCrtPath string) *MyHttpClientBuilder {
	b.VerifySslClientCrtPath = verifySslClientCrtPath
	return b
}

func (b *MyHttpClientBuilder) SetVerifySslClientKeyPath(verifySslClientKeyPath string) *MyHttpClientBuilder {
	b.VerifySslClientKeyPath = verifySslClientKeyPath
	return b
}

func (b *MyHttpClientBuilder) SetVerifySslCaCrtPath(verifySslCaCrtPath string) *MyHttpClientBuilder {
	b.VerifySslCaCrtPath = verifySslCaCrtPath
	return b
}

func (b *MyHttpClientBuilder) Build() *MyHttpClient {
	return &MyHttpClient{
		Client:  http.DefaultClient,
		Builder: b,
	}
}
