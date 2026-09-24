package data

import (
	"context"
	_ "encoding/json"
	"errors"
)

type PlayerItem struct {
	PlayerPubID string `json:"player-id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Job         string `json:"job"`
	Description string `json:"description"`
	AreaCode    string `json:"area-code"`
	AreaName    string `json:"area-name"`
	ImagePubID  string `json:"image-id"`
	Free        bool   `json:"free"`
}

func PlayerList(regionCode string) ([]PlayerItem, error) {
	rows, err := db.Query(context.Background(), `
		SELECT p.pub_id, f.name, f.race, f.job, f.description,
		  f.area_code, f.area_name, f.image_pub_id,
		  p.user_id IS NULL AS free
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
			&item.PlayerPubID,
			&item.Name,
			&item.Race,
			&item.Job,
			&item.Description,
			&item.AreaCode,
			&item.AreaName,
			&item.ImagePubID,
			&item.Free,
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
	ImagePubID  string `json:"image-id"`
}

const actionsPerPage = 12

func ActionList(playerPubID string, page int) ([]ActionItem, bool, error) {
	if page < 1 {
		return nil, false, errors.New("invalid page")
	}

	offset := (page - 1) * actionsPerPage

	rows, err := db.Query(context.Background(), `
		SELECT a.day, a.description, a.image_pub_id
		FROM actions AS a
		JOIN flavors AS f ON f.id = a.flavor_id
		JOIN players AS p ON p.id = f.player_id
		WHERE p.pub_id = $1 AND a.fixed = TRUE
		ORDER BY a.id DESC
		LIMIT $2 OFFSET $3
	`, playerPubID, actionsPerPage+1, offset)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	var list []ActionItem

	for rows.Next() {
		var item ActionItem
		var day int64

		err := rows.Scan(
			&day,
			&item.Description,
			&item.ImagePubID,
		)
		if err != nil {
			return nil, false, err
		}

		date := NewDate(day)
		item.Date = date.String()

		list = append(list, item)
	}

	if err := rows.Err(); err != nil {
		return nil, false, err
	}

	hasNext := len(list) > actionsPerPage
	if hasNext {
		list = list[:actionsPerPage]
	}

	return list, hasNext, nil
}

func PlayerCounts() (map[string]int, error) {
	rows, err := db.Query(context.Background(), `
		SELECT
		  split_part(f.area_code, '-', 1) AS region_code,
		  COUNT(*) AS player_count
		FROM players AS p
		JOIN LATERAL (
		  SELECT area_code
		  FROM flavors
		  WHERE player_id = p.id
		    AND committed = TRUE
		  ORDER BY id DESC
		  LIMIT 1
		) AS f ON TRUE
		WHERE p.activated = TRUE
		GROUP BY region_code
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[string]int)

	for rows.Next() {
		var regionCode string
		var count int

		if err := rows.Scan(&regionCode, &count); err != nil {
			return nil, err
		}

		counts[regionCode] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return counts, nil
}

func RegionStates() (map[string]string, error) {
	rows, err := db.Query(context.Background(), `
		SELECT DISTINCT ON (region_code)
		  region_code, state
		FROM region_states
		ORDER BY region_code, day DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	states := make(map[string]string)

	for rows.Next() {
		var regionCode string
		var state string

		if err := rows.Scan(&regionCode, &state); err != nil {
			return nil, err
		}

		states[regionCode] = state
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return states, nil
}
