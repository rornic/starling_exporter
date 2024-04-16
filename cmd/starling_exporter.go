package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rornic/starlingexporter/internal/client"
	"github.com/rornic/starlingexporter/internal/metrics"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	client := client.NewStarlingHttpClient()
	metrics.Record(&client)

	http.Handle("/metrics", promhttp.Handler())

	slog.Info("listening on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("error running server: %v\n", err)
		os.Exit(1)
	}
}
