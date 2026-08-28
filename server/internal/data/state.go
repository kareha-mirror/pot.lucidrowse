package data

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func AreaState(areaCode string) (string, error) {
	var state string

	err := db.QueryRow(context.Background(), `
		SELECT state FROM area_states
		WHERE area_code = $1
		ORDER BY day DESC
		LIMIT 1
		`, areaCode).Scan(&state)

	if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return state, nil
}

func AddAreaState(areaCode string, areaState string) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO area_states (area_code, day, state)
		SELECT $1, day_counter, $2
		FROM state
		WHERE id = 1
		`, areaCode, areaState)
	if err != nil {
		return err
	}

	return nil
}
