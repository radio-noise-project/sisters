package main

import (
	"log/slog"

	"github.com/radio-noise-project/sisters/internal/api"
)

func main() {
	slog.Info("Start sisters")
	api.Server()
}
