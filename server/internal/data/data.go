package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

var db *pgxpool.Pool

const createTablesSql = `
CREATE TABLE IF NOT EXISTS state (
	id INT PRIMARY KEY DEFAULT 1,
	boot_counter BIGINT NOT NULL DEFAULT 0,
	day_counter BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS players (
	id SERIAL PRIMARY KEY,
	pub_id TEXT NOT NULL UNIQUE,
	day BIGINT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	activated BOOL NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS images (
	pub_id TEXT PRIMARY KEY,
	content_type TEXT NOT NULL,
	content BYTEA NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS flavors (
	id SERIAL PRIMARY KEY,
	player_id INT NOT NULL REFERENCES players(id),
	input TEXT NOT NULL,
	name TEXT NOT NULL,
	race TEXT NOT NULL,
	job TEXT NOT NULL,
	description TEXT NOT NULL,
	area_code TEXT NOT NULL,
	area_name TEXT NOT NULL,
	day BIGINT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	image_pub_id TEXT REFERENCES images(pub_id),
	committed BOOL NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS actions (
	id SERIAL PRIMARY KEY,
	flavor_id INT NOT NULL REFERENCES flavors(id),
	input TEXT NOT NULL,
	description TEXT NOT NULL,
	day BIGINT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	image_pub_id TEXT REFERENCES images(pub_id),
	committed BOOL NOT NULL DEFAULT FALSE,
	fixed BOOL NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS areas (
	id SERIAL PRIMARY KEY,
	region_code TEXT NOT NULL,
	area_code TEXT NOT NULL,
	name TEXT NOT NULL,

	UNIQUE (region_code, area_code)
);

CREATE TABLE IF NOT EXISTS area_states (
	id SERIAL PRIMARY KEY,
	area_code TEXT NOT NULL,
	day BIGINT NOT NULL,
	state TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),

	UNIQUE (area_code, day)
);

CREATE TABLE IF NOT EXISTS region_states (
	id SERIAL PRIMARY KEY,
	region_code TEXT NOT NULL,
	day BIGINT NOT NULL,
	state TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),

	UNIQUE (region_code, day)
);

CREATE TABLE IF NOT EXISTS world_states (
	id SERIAL PRIMARY KEY,
	day BIGINT NOT NULL UNIQUE,
	state TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now()
);
`

var bootCount int64

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	var err error
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(context.Background(), createTablesSql)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(context.Background(), `
		INSERT INTO state (id, boot_counter, day_counter)
		VALUES (1, 1, 0)
		ON CONFLICT (id)
		DO UPDATE SET boot_counter = state.boot_counter + 1
		RETURNING boot_counter;
		`).Scan(&bootCount)
	if err != nil {
		return nil, err
	}

	return db, nil
}
