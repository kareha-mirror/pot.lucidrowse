package main

import (
	"os"

	"tea.kareha.org/pot/lucidrowse/server/internal/api"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		print("Usage: " + os.Args[0] + " server.yaml\n")
		return
	}
	cfg := config.Load(os.Args[1])
	api.Run(cfg)
}
