package main

import (
	"fmt"
	"os"

	"tea.kareha.org/pot/lucidrowse/server/internal/api"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

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

	if err = api.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
}
