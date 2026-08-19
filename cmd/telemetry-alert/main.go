package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/example/telemetry-alert/internal/app"
	"github.com/example/telemetry-alert/internal/config"
	"github.com/example/telemetry-alert/internal/logging"
)

func main() {
	configPath := flag.String("config", "configs/config.yaml", "path to YAML config file")
	flag.Parse()

	logger := logging.New("info")
	cfg, err := config.Load(*configPath)
	if err != nil {
		logger.Error("load config", "error", err)
		os.Exit(1)
	}
	logger = logging.New(cfg.Logging.Level)

	application, err := app.New(cfg, logger)
	if err != nil {
		logger.Error("initialize application", "error", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if err := application.Run(ctx); err != nil {
		logger.Error("application stopped with error", "error", err)
		os.Exit(1)
	}
	logger.Info("application stopped")
}
