package data

import (
	"context"
	_ "encoding/json"
)

type PlayerItem struct {
	PlayerPubId string `json:"player-id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Job         string `json:"job"`
	Description string `json:"description"`
	ImagePubId  string `json:"image-id"`
}

func PlayerList() ([]PlayerItem, error) {
	rows, err := db.Query(context.Background(), `
		SELECT p.pub_id, f.name, f.race, f.job, f.description, f.image_pub_id
		FROM players AS p
		JOIN LATERAL (
		  SELECT name, race, job, description, image_pub_id
		  FROM flavors
		  WHERE player_id = p.id AND committed = TRUE
		  ORDER BY id DESC
		  LIMIT 1
		) AS f ON TRUE
		WHERE p.activated = TRUE
		ORDER BY p.id
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []PlayerItem

	for rows.Next() {
		var p PlayerItem

		err := rows.Scan(
			&p.PlayerPubId,
			&p.Name,
			&p.Race,
			&p.Job,
			&p.Description,
			&p.ImagePubId,
		)
		if err != nil {
			return nil, err
		}

		players = append(players, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

type ActionItem struct {
	Day         int64  `json:"day"`
	Description string `json:"description"`
	ImagePubId  string `json:"image-id"`
}

func ActionList(playerPubId string) ([]ActionItem, error) {
	rows, err := db.Query(context.Background(), `
		SELECT a.day, a.description, a.image_pub_id
		FROM actions AS a
		JOIN players AS p ON p.id = a.player_id
		WHERE p.pub_id = $1 AND a.fixed = TRUE
		ORDER BY a.id DESC
		`, playerPubId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []ActionItem

	for rows.Next() {
		var a ActionItem

		err := rows.Scan(
			&a.Day,
			&a.Description,
			&a.ImagePubId,
		)
		if err != nil {
			return nil, err
		}

		actions = append(actions, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actions, nil
}
