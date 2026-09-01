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
	AreaCode    string `json:"area-code"`
	AreaName    string `json:"area-name"`
	ImagePubId  string `json:"image-id"`
}

func PlayerList(regionCode string) ([]PlayerItem, error) {
	rows, err := db.Query(context.Background(), `
		SELECT p.pub_id, f.name, f.race, f.job, f.description,
		  f.area_code, f.area_name, f.image_pub_id
		FROM players AS p
		JOIN LATERAL (
		  SELECT name, race, job, description,
		    area_code, area_name, image_pub_id
		  FROM flavors
		  WHERE player_id = p.id AND committed = TRUE
		  ORDER BY id DESC
		  LIMIT 1
		) AS f ON TRUE
		WHERE p.activated = TRUE
		  AND f.area_code LIKE $1 || '%'
		ORDER BY p.id
	`, regionCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []PlayerItem

	for rows.Next() {
		var item PlayerItem

		err := rows.Scan(
			&item.PlayerPubId,
			&item.Name,
			&item.Race,
			&item.Job,
			&item.Description,
			&item.AreaCode,
			&item.AreaName,
			&item.ImagePubId,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

type ActionItem struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	ImagePubId  string `json:"image-id"`
}

func ActionList(playerPubId string) ([]ActionItem, error) {
	rows, err := db.Query(context.Background(), `
		SELECT a.day, a.description, a.image_pub_id
		FROM actions AS a
		JOIN flavors AS f ON f.id = a.flavor_id
		JOIN players AS p ON p.id = f.player_id
		WHERE p.pub_id = $1 AND a.fixed = TRUE
		ORDER BY a.id DESC
	`, playerPubId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []ActionItem

	for rows.Next() {
		var item ActionItem
		var day int64

		err := rows.Scan(
			&day,
			&item.Description,
			&item.ImagePubId,
		)
		if err != nil {
			return nil, err
		}

		date := NewDate(day)
		item.Date = date.String()

		list = append(list, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
