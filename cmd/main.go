package main

import (
	"log/slog"
	"os"
	"reminder_bot/internal/app"
	"reminder_bot/internal/config"
)

func main() {
	cfg := config.NewConfig()
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	application := app.New(log, cfg.APIToken, "fef")

	app.Run(application)
}
