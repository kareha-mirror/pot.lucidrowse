package main

import (
	"fmt"
	"os"

	"tea.kareha.org/pot/lucidrowse/server/internal/api"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s server.yaml\n", os.Args[0])
		return
	}

	cfg := config.Load(os.Args[1])

	api.Run(cfg)
}
