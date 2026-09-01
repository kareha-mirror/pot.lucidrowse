package data

import "context"

func SeedWorld(worldStateSeed string) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO world_states (day, state)
		SELECT day_counter, $1
		FROM state
		WHERE id = 1
		  AND NOT EXISTS (
		  	SELECT 1 FROM world_states
		  )
	`, worldStateSeed)

	return err
}
