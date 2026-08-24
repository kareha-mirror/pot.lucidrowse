package data

import "context"

func SaveImage(pubId string, content []byte) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO images (pub_id, content) VALUES ($1, $2)
		`, pubId, content)
	if err != nil {
		return err
	}
	return nil
}

func LoadImage(pubId string) ([]byte, error) {
	var content []byte
	err := db.QueryRow(context.Background(), `
		SELECT content FROM images WHERE pub_id=$1
		`, pubId,
	).Scan(&content)
	if err != nil {
		return nil, err
	}
	return content, nil
}
