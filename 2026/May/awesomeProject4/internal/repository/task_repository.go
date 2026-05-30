package repository

import (
	"context"
	"errors"
	"fmt"

	"app/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}

func (r *TaskRepository) List(ctx context.Context, userID int) ([]models.Task, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, title, description, created_at, updated_at
		FROM tasks
		WHERE user_id = $1
		ORDER BY id DESC`, userID)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func (r *TaskRepository) GetByID(ctx context.Context, userID, id int) (*models.Task, error) {
	var t models.Task
	err := r.pool.QueryRow(ctx, `
		SELECT id, title, description, created_at, updated_at
		FROM tasks WHERE id = $1 AND user_id = $2`, id, userID).
		Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get task: %w", err)
	}
	return &t, nil
}

func (r *TaskRepository) Create(ctx context.Context, userID int, title, description string) (*models.Task, error) {
	var t models.Task
	err := r.pool.QueryRow(ctx, `
		INSERT INTO tasks (user_id, title, description)
		VALUES ($1, $2, $3)
		RETURNING id, title, description, created_at, updated_at`,
		userID, title, description).
		Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}
	return &t, nil
}

func (r *TaskRepository) Update(ctx context.Context, userID, id int, title, description string) (*models.Task, error) {
	var t models.Task
	err := r.pool.QueryRow(ctx, `
		UPDATE tasks
		SET title = $1, description = $2, updated_at = NOW()
		WHERE id = $3 AND user_id = $4
		RETURNING id, title, description, created_at, updated_at`,
		title, description, id, userID).
		Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("update task: %w", err)
	}
	return &t, nil
}

func (r *TaskRepository) Delete(ctx context.Context, userID, id int) error {
	tag, err := r.pool.Exec(ctx, `DELETE FROM tasks WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
