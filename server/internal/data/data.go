package data

import (
	"context"
	"fmt"
	"log"

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
`

var bootCount int64

func Connect(cfg *config.Config) *pgxpool.Pool {
	var err error

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_, err = db.Exec(context.Background(), createTablesSql)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	err = db.QueryRow(context.Background(), `
		INSERT INTO state (id, boot_counter, day_counter)
		VALUES (1, 1, 0)
		ON CONFLICT (id)
		DO UPDATE SET boot_counter = state.boot_counter + 1
		RETURNING boot_counter;
	`).Scan(&bootCount)
	if err != nil {
		log.Fatalf("failed to count up boot conter: %v", err)
	}

	return db
}

func Day() (int64, error) {
	var dayCount int64
	err := db.QueryRow(
		context.Background(),
		"SELECT day_counter FROM state WHERE id = 1",
	).Scan(&dayCount)
	if err != nil {
		return 0, err
	}

	return dayCount, nil
}

func NextDay() (int64, error) {
	ctx := context.Background()

	var dayCount int64
	err := db.QueryRow(ctx, `
		UPDATE state
		SET day_counter = day_counter + 1
		WHERE id = 1
		RETURNING day_counter
	`).Scan(&dayCount)
	if err != nil {
		return 0, err
	}

	return dayCount, nil
}
