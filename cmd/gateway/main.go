package main

import (
	"flag"
	"fmt"
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

	p := proxy.New(cfg)
	slog.Info("Flux Gateway running", slog.Int("port", cfg.Server.Port))

	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := http.ListenAndServe(serverAddr, p); err != nil {
		slog.Error("Gateway server crashed", slog.Any("error", err))
		os.Exit(1)
	}
}
