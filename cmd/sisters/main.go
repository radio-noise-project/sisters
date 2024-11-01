package main

import (
	"log/slog"

	"github.com/radio-noise-project/sisters/internal/pkg/container"
)

func main() {
	slog.Info("Start sisters")

	// Start a container
	var status = container.ContainerCreateAndStart()
	slog.Info(string(status))
}
