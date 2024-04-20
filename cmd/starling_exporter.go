package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rornic/starlingexporter/internal/client"
	"github.com/rornic/starlingexporter/internal/metrics"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	accessToken := os.Getenv("STARLING_ACCESS_TOKEN")
	if accessToken == "" {
		slog.Error("STARLING_ACCESS_TOKEN is not set. Exiting.")
		os.Exit(1)
	}
	slog.Info("using access token from environment")

	endpoint := "https://api.starlingbank.com/api/v2"
	sandbox := strings.ToLower(os.Getenv("STARLING_SANDBOX")) == "true"
	if sandbox {
		slog.Info("using sandbox environment")
		endpoint = strings.Replace(endpoint, "api", "api-sandbox", 1)
	}

	client := client.NewStarlingHttpClient(accessToken, endpoint)
	metrics.Record(&client)

	http.Handle("/metrics", promhttp.Handler())

	slog.Info("listening on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("error running server: %v\n", err)
		os.Exit(1)
	}
}
