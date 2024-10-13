package main

import (
	"log/slog"
	"os"
	"reminder_bot/internal/app"
	"reminder_bot/internal/config"
	"reminder_bot/internal/database"
)

func main() {
	cfg := config.NewConfig()
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	db := database.Connect(cfg.DBPath)

	application := app.New(log, cfg.APIToken)

	app.Run(application)
}
