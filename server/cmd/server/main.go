package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"tea.kareha.org/pot/lucidrowse/server/internal/api"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func StartDayTicker(cfg *config.Config, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			if err := api.NextDay(cfg); err != nil {
				log.Printf("failed to go next day: %v", err)
			}
		}
	}()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s server.yaml\n", os.Args[0])
		return
	}

	cfg, err := config.Load(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	db, err := data.Connect(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	defer db.Close()

	if err = data.Seed(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	interval, err := time.ParseDuration(cfg.Game.DayDuration)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	StartDayTicker(cfg, interval)

	if err = api.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
}
