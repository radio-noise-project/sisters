package main

import (
	"log/slog"

	"github.com/radio-noise-project/sisters/internal/api/server"
)

func main() {
	slog.Info("Start sisters")
	server.Start()
}
