package data

import "context"

type Image struct {
	ContentType string
	Content     []byte
}

func SaveImage(ctx context.Context, pubID string, image Image) error {
	_, err := db.Exec(ctx, `
		INSERT INTO images (pub_id, content_type, content)
		VALUES ($1, $2, $3)
	`, pubID, image.ContentType, image.Content)
	if err != nil {
		return err
	}
	return nil
}

func LoadImage(ctx context.Context, pubID string) (Image, error) {
	var image Image
	err := db.QueryRow(ctx, `
		SELECT content_type, content FROM images WHERE pub_id = $1
	`, pubID).Scan(
		&image.ContentType, &image.Content,
	)
	if err != nil {
		return Image{}, err
	}
	return image, nil
}
