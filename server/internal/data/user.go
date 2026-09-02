package data

import "context"

func CreateUser() (int, error) {
	var userID int
	err := db.QueryRow(context.Background(), `
		INSERT INTO users DEFAULT VALUES
		RETURNING id
	`).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func AddSession(userID int, keyHash []byte) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO sessions (user_id, key_hash)
		VALUES ($1, $2)
	`, userID, keyHash)
	if err != nil {
		return err
	}
	return nil
}

func UserID(keyHash []byte) (int, error) {
	var userID int
	err := db.QueryRow(context.Background(), `
		SELECT user_id FROM sessions WHERE key_hash = $1
	`, keyHash).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
