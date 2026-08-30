package data

import (
	"context"
	_ "encoding/json"
)

type EventItem struct {
	Name        string `json:"name"`
	Race        string `json:"race"`
	Job         string `json:"job"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

func EventList(areaCode string) ([]EventItem, error) {
	rows, err := db.Query(context.Background(), `
		SELECT f.name, f.race, f.job, f.description, a.description
		FROM actions AS a
		JOIN flavors AS f ON f.id = a.flavor_id
		WHERE a.committed = TRUE
		  AND a.day = (SELECT day_counter FROM state WHERE id = 1)
		  AND f.area_code = $1
		ORDER BY a.id ASC
	`, areaCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []EventItem

	for rows.Next() {
		var item EventItem

		err := rows.Scan(
			&item.Name,
			&item.Race,
			&item.Job,
			&item.Description,
			&item.Action,
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
