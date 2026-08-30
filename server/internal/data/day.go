package data

import "context"

func Day() (int64, error) {
	var day int64
	err := db.QueryRow(context.Background(), `
		SELECT day_counter FROM state WHERE id = 1
	`).Scan(&day)
	if err != nil {
		return 0, err
	}
	return day, nil
}

func NextDay() (int64, error) {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		UPDATE actions
		SET fixed = TRUE
		WHERE day = (SELECT day_counter FROM state WHERE id = 1)
		AND committed = TRUE
		AND fixed = FALSE
	`)
	if err != nil {
		return 0, err
	}

	var day int64
	err = tx.QueryRow(ctx, `
		UPDATE state
		SET day_counter = day_counter + 1
		WHERE id = 1
		RETURNING day_counter
	`).Scan(&day)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return day, nil
}
