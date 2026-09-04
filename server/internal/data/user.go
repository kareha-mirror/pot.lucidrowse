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

type User struct {
	ID      int
	Name    *string
	AICalls int
}

func LoadUser(keyHash []byte) (User, error) {
	var user User
	err := db.QueryRow(context.Background(), `
		SELECT s.user_id, u.name, u.ai_calls
		FROM sessions AS s
		JOIN users AS u ON u.id = s.user_id
		WHERE s.key_hash = $1
	`, keyHash).Scan(&user.ID, &user.Name, &user.AICalls)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func IncrementAICalls(userID int) error {
	_, err := db.Exec(context.Background(), `
		UPDATE users
		SET ai_calls = ai_calls + 1
		WHERE id = $1
	`, userID)
	return err
}
