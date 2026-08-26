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
		JOIN players AS p
		  ON p.id = a.player_id
		JOIN LATERAL (
		  SELECT name, race, job, description, area_code
		  FROM flavors
		  WHERE player_id = p.id
		    AND committed = TRUE
		  ORDER BY id DESC
		  LIMIT 1
		) AS f ON TRUE
		WHERE a.committed = TRUE
		  AND a.day = (SELECT day_counter FROM state WHERE id = 1)
		  AND f.area_code = $1
		ORDER BY a.id ASC;
		`, areaCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []EventItem

	for rows.Next() {
		var i EventItem

		err := rows.Scan(
			&i.Name,
			&i.Race,
			&i.Job,
			&i.Description,
			&i.Action,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, i)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
