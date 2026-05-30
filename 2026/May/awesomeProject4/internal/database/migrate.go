package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(ctx context.Context, pool *pgxpool.Pool) error {
	if _, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id            SERIAL PRIMARY KEY,
			username      VARCHAR(32) NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`); err != nil {
		return fmt.Errorf("migrate users table: %w", err)
	}

	if _, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS tasks (
			id          SERIAL PRIMARY KEY,
			user_id     INTEGER REFERENCES users(id) ON DELETE CASCADE,
			title       VARCHAR(255) NOT NULL,
			description TEXT NOT NULL DEFAULT '',
			created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`); err != nil {
		return fmt.Errorf("migrate tasks table: %w", err)
	}

	if _, err := pool.Exec(ctx, `
		ALTER TABLE tasks ADD COLUMN IF NOT EXISTS user_id INTEGER REFERENCES users(id) ON DELETE CASCADE`); err != nil {
		return fmt.Errorf("migrate tasks user_id: %w", err)
	}

	if _, err := pool.Exec(ctx, `
		CREATE INDEX IF NOT EXISTS idx_tasks_user_id ON tasks(user_id)`); err != nil {
		return fmt.Errorf("migrate tasks index: %w", err)
	}

	return nil
}
