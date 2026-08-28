package data

import "context"

func CreatePlayer(playerPubId string) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO players
		(pub_id, day)
		SELECT $1, day_counter
		FROM state
		WHERE id = 1
		`, playerPubId)
	if err != nil {
		return err
	}
	return nil
}

func PlayerId(playerPubId string) (int, error) {
	var playerId int
	err := db.QueryRow(context.Background(), `
		SELECT id FROM players WHERE pub_id=$1
		`, playerPubId).
		Scan(&playerId)
	if err != nil {
		return 0, err
	}
	return playerId, nil
}

func ActivatePlayer(playerId int) error {
	_, err := db.Exec(context.Background(), `
		UPDATE players SET activated=TRUE WHERE id=$1
		`, playerId)
	return err
}

type Flavor struct {
	Input       string
	Name        string
	Race        string
	Job         string
	Description string
	AreaCode    string
	AreaName    string

	// optionals
	Day        int64
	ImagePubId *string
	Committed  bool
}

func AddFlavor(playerId int, flavor Flavor) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO flavors
		(player_id, input, name, race, job, description, area_code, area_name, day)
		SELECT $1, $2, $3, $4, $5, $6, $7, $8, day_counter
		FROM state
		WHERE id = 1
		`, playerId, flavor.Input, flavor.Name, flavor.Race, flavor.Job, flavor.Description, flavor.AreaCode, flavor.AreaName)
	if err != nil {
		return err
	}
	return nil
}

func AddImageToLastFlavor(playerId int, imagePubId string) error {
	_, err := db.Exec(context.Background(), `
		UPDATE flavors
		SET image_pub_id=$1
		WHERE id = (
		  SELECT id
		  FROM flavors
		  WHERE player_id=$2
		  ORDER BY id DESC
		  LIMIT 1
		)
		`, imagePubId, playerId)
	return err
}

func CommitLastFlavor(playerId int) error {
	_, err := db.Exec(context.Background(), `
		UPDATE flavors
		SET committed=TRUE
		WHERE id = (
		  SELECT id
		  FROM flavors
		  WHERE player_id=$1
		  ORDER BY id DESC
		  LIMIT 1
		)
		`, playerId)
	return err
}

func LoadLastFlavor(playerId int) (Flavor, error) {
	var flavor Flavor
	err := db.QueryRow(context.Background(), `
		SELECT name, race, job, description, area_code, area_name, day, image_pub_id, committed
		FROM flavors
		WHERE id = (
		  SELECT id
		  FROM flavors
		  WHERE player_id=$1
		  ORDER BY id DESC
		  LIMIT 1
		)
		`, playerId).
		Scan(&flavor.Name, &flavor.Race, &flavor.Job, &flavor.Description, &flavor.AreaCode, &flavor.AreaName, &flavor.Day, &flavor.ImagePubId, &flavor.Committed)
	if err != nil {
		return Flavor{}, err
	}
	return flavor, nil
}

type Action struct {
	Input       string
	Description string

	// optional
	Day        int64
	ImagePubId *string
	Committed  bool
	Fixed      bool
}

func AddAction(playerId int, action Action) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO actions
		(flavor_id, input, description, day)
		SELECT f.id, $2, $3, day_counter
		FROM flavors f
		CROSS JOIN state s
		WHERE f.player_id = $1
		  AND f.committed = TRUE
		  AND s.id = 1
		ORDER BY f.id DESC
		LIMIT 1
		`, playerId, action.Input, action.Description)
	if err != nil {
		return err
	}
	return nil
}

func AddImageToLastAction(playerId int, imagePubId string) error {
	_, err := db.Exec(context.Background(), `
		UPDATE actions
		SET image_pub_id=$1
		WHERE id = (
		  SELECT a.id
		  FROM actions a
		  JOIN flavors f ON a.flavor_id = f.id
		  WHERE f.player_id = $2
		  ORDER BY a.id DESC
		  LIMIT 1
		)
		`, imagePubId, playerId)
	return err
}

func CommitLastAction(playerId int) error {
	_, err := db.Exec(context.Background(), `
		UPDATE actions
		SET committed = TRUE
		WHERE id = (
		  SELECT a.id
		  FROM actions a
		  JOIN flavors f ON a.flavor_id = f.id
		  WHERE f.player_id = $1
		  ORDER BY a.id DESC
		  LIMIT 1
		)
		`, playerId)
	return err
}

func LoadLastAction(playerId int) (Action, error) {
	var action Action
	err := db.QueryRow(context.Background(), `
		SELECT description, day, image_pub_id, committed, fixed
		FROM actions
		WHERE id = (
		  SELECT a.id
		  FROM actions a
		  JOIN flavors f ON a.flavor_id = f.id
		  WHERE f.player_id = $1
		  ORDER BY a.id DESC
		  LIMIT 1
		)
		`, playerId).
		Scan(&action.Description, &action.Day, &action.ImagePubId, &action.Committed, &action.Fixed)
	if err != nil {
		return Action{}, err
	}
	return action, nil
}
