package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/faizkhan-06/flux-gateway/internal/config"
	"github.com/faizkhan-06/flux-gateway/internal/logger"
	"github.com/faizkhan-06/flux-gateway/internal/proxy"
)

func main() {
	filepathFlag := flag.String("config", "routes.yaml", "Path to your routing configuration file")
	flag.Parse()

	logger.Init()

	cfg, err := config.Load(*filepathFlag)
	if err != nil {
		slog.Error("Failed to load gateway configuration", slog.String("path", *filepathFlag), slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("Configuration loaded successfully",
		slog.Int("server_port", cfg.Server.Port),
		slog.Int("total_routes", len(cfg.Routes)),
	)

	hardcodedTarget := "http://localhost:8081"
	p := proxy.New(hardcodedTarget)
	slog.Info("Flux Gateway running", slog.String("port", "8080"), slog.String("target", hardcodedTarget))

	if err := http.ListenAndServe(":8080", p); err != nil {
		slog.Error("Gateway server crashed", slog.Any("error", err))
		os.Exit(1)
	}
}
