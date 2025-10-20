package config

import (
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
)

func ViperConfig() {
	path := AddConfigPath()
	if !lib.FileExists(path + "/config.yaml") {
		log.Fatalf("Error in no config yaml found.")
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("")

	err := viper.ReadInConfig()
	if err != nil {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Error("Error reading the config file", "detail", err)
		os.Exit(1)
		return
	}
	viper.AutomaticEnv()

	err = viper.MergeInConfig()
	if err != nil {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Error("Error reading the static_responses file", "detail", err)
		os.Exit(1)
		return
	}

}

func AddConfigPath() string {
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Error getting the current dir", "detail", err)
		os.Exit(1)
	}
	dir, _ = strings.CutPrefix(dir, "/")
	if strings.HasSuffix(dir, "cmd/app") || strings.HasSuffix(dir, "cmd/test") || strings.HasSuffix(dir, "bin") {
		path := "../.."
		viper.AddConfigPath(path)
		return path
	}
	path := "."
	viper.AddConfigPath(path)
	return path
}
