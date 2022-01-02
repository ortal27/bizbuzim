// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package dal

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createExpense = `-- name: CreateExpense :exec
INSERT INTO expenses (id, name, payment, price, tags, description, "createdAt", "createdBy")
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreateExpenseParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Payment     string    `json:"payment"`
	Price       string    `json:"price"`
	Tags        []string  `json:"tags"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   string    `json:"createdBy"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) error {
	_, err := q.db.ExecContext(ctx, createExpense,
		arg.ID,
		arg.Name,
		arg.Payment,
		arg.Price,
		pq.Array(arg.Tags),
		arg.Description,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	return err
}

const createRawExpense = `-- name: CreateRawExpense :exec
INSERT INTO raw_expenses (id, text, "createdAt", "createdBy")
    VALUES ($1, $2, $3, $4)
`

type CreateRawExpenseParams struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}

func (q *Queries) CreateRawExpense(ctx context.Context, arg CreateRawExpenseParams) error {
	_, err := q.db.ExecContext(ctx, createRawExpense,
		arg.ID,
		arg.Text,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	return err
}