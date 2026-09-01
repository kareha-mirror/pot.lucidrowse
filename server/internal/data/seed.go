package data

import "tea.kareha.org/pot/lucidrowse/server/internal/config"

func Seed(cfg *config.Config) error {
	/*
		err := SeedWorld(cfg.World.StateSeed)
		if err != nil {
			return err
		}
	*/

	err := SeedAreas()
	if err != nil {
		return err
	}

	return nil
}
