package data

import "context"

func CreatePlayer(userID int, playerPubID string) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO players (user_id, pub_id, day)
		SELECT $1, $2, day_counter
		FROM state WHERE id = 1
	`, userID, playerPubID)
	if err != nil {
		return err
	}
	return nil
}

type Player struct {
	ID        int
	PubID     string
	Day       int64
	Activated bool
}

func LoadPlayer(userID int) (Player, error) {
	var player Player
	err := db.QueryRow(context.Background(), `
		SELECT id, pub_id, day, activated
		FROM players
		WHERE user_id = $1
	`, userID).Scan(
		&player.ID, &player.PubID, &player.Day, &player.Activated,
	)
	if err != nil {
		return Player{}, err
	}
	return player, nil
}

func ActivatePlayer(playerID int) error {
	_, err := db.Exec(context.Background(), `
		UPDATE players SET activated = TRUE WHERE id = $1
	`, playerID)
	return err
}

func ReleasePlayer(userID int) error {
	_, err := db.Exec(context.Background(), `
		UPDATE players
		SET user_id = NULL
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return err
	}
	return nil
}

func OverridePlayer(userID int, playerPubID string) error {
	_, err := db.Exec(context.Background(), `
		UPDATE players
		SET user_id = $1
		WHERE pub_id = $2 AND user_id IS NULL
	`, userID, playerPubID)
	if err != nil {
		return err
	}
	return nil
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
	ImagePubID *string
	Committed  bool
}

func AddFlavor(playerID int, flavor Flavor) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO flavors
		(player_id, input, name, race, job, description,
		  area_code, area_name, day)
		SELECT $1, $2, $3, $4, $5, $6, $7, $8, day_counter
		FROM state WHERE id = 1
	`, playerID, flavor.Input, flavor.Name, flavor.Race, flavor.Job, flavor.Description, flavor.AreaCode, flavor.AreaName)
	if err != nil {
		return err
	}
	return nil
}

func AddImageToLastFlavor(playerID int, imagePubID string) error {
	_, err := db.Exec(context.Background(), `
		UPDATE flavors
		SET image_pub_id = $1
		WHERE id = (
		  SELECT id
		  FROM flavors
		  WHERE player_id = $2
		  ORDER BY id DESC
		  LIMIT 1
		)
	`, imagePubID, playerID)
	return err
}

func CommitLastFlavor(playerID int) error {
	_, err := db.Exec(context.Background(), `
		UPDATE flavors
		SET committed=TRUE
		WHERE id = (
		  SELECT id
		  FROM flavors
		  WHERE player_id = $1
		  ORDER BY id DESC
		  LIMIT 1
		)
	`, playerID)
	return err
}

func LoadLastFlavor(playerID int) (Flavor, error) {
	var flavor Flavor
	err := db.QueryRow(context.Background(), `
		SELECT name, race, job, description,
		  area_code, area_name, day, image_pub_id, committed
		FROM flavors
		WHERE id = (
		  SELECT id
		  FROM flavors
		  WHERE player_id = $1
		  ORDER BY id DESC
		  LIMIT 1
		)
	`, playerID).Scan(
		&flavor.Name, &flavor.Race, &flavor.Job, &flavor.Description,
		&flavor.AreaCode, &flavor.AreaName, &flavor.Day,
		&flavor.ImagePubID, &flavor.Committed,
	)
	if err != nil {
		return Flavor{}, err
	}
	return flavor, nil
}

func LoadCurrentFlavor(playerID int) (Flavor, error) {
	var flavor Flavor
	err := db.QueryRow(context.Background(), `
		SELECT name, race, job, description,
		  area_code, area_name, day, image_pub_id, committed
		FROM flavors
		WHERE id = (
		  SELECT id
		  FROM flavors
		  WHERE player_id = $1 AND committed = TRUE
		  ORDER BY id DESC
		  LIMIT 1
		)
	`, playerID).Scan(
		&flavor.Name, &flavor.Race, &flavor.Job, &flavor.Description,
		&flavor.AreaCode, &flavor.AreaName, &flavor.Day,
		&flavor.ImagePubID, &flavor.Committed,
	)
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
	ImagePubID *string
	Committed  bool
	Fixed      bool
}

func AddAction(playerID int, action Action) error {
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
	`, playerID, action.Input, action.Description)
	if err != nil {
		return err
	}
	return nil
}

func AddImageToLastAction(playerID int, imagePubID string) error {
	_, err := db.Exec(context.Background(), `
		UPDATE actions
		SET image_pub_id = $1
		WHERE id = (
		  SELECT a.id
		  FROM actions a
		  JOIN flavors f ON a.flavor_id = f.id
		  WHERE f.player_id = $2
		  ORDER BY a.id DESC
		  LIMIT 1
		)
	`, imagePubID, playerID)
	return err
}

func CommitLastAction(playerID int) error {
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
	`, playerID)
	return err
}

func LoadCurrentAction(playerID int) (Action, error) {
	var action Action
	err := db.QueryRow(context.Background(), `
		SELECT a.description, a.day, a.image_pub_id, a.committed, a.fixed
		FROM actions a
		JOIN flavors f ON a.flavor_id = f.id
		WHERE f.player_id = $1
		  AND a.committed = TRUE
		  AND a.day = (SELECT day_counter FROM state WHERE id = 1)
		ORDER BY a.id DESC
		LIMIT 1
	`, playerID).Scan(
		&action.Description, &action.Day, &action.ImagePubID,
		&action.Committed, &action.Fixed,
	)
	if err != nil {
		return Action{}, err
	}
	return action, nil
}

func LoadLastAction(playerID int) (Action, error) {
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
	`, playerID).Scan(
		&action.Description, &action.Day, &action.ImagePubID,
		&action.Committed, &action.Fixed,
	)
	if err != nil {
		return Action{}, err
	}
	return action, nil
}
