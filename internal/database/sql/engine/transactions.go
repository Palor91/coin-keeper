package engine

import (
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (e *DatabaseEngine) CreateTransaction(ctx context.Context, transaction write.Transactions) error {
	_, err := e.db.ExecContext(
		ctx,
		"INSERT INTO transactions (id, user_id, description, amount, date) VALUES ($1, $2, $3, $4, $5)",
		transaction.ID, transaction.User_id, transaction.Description, transaction.Amount, transaction.Date,
	)
	if err != nil {
		log.Error().Msgf("Wrong creation: %v", err)
	}
	return err
}

func (e *DatabaseEngine) ReadTransaction(ctx context.Context, transactionID int) (read.Transactions, error) {
	rows, err := e.db.QueryContext(
		ctx,
		"SELECT (id, user_id, description, amount, date) FROM users WHERE id = $1",
		transactionID,
	)
	if err != nil {
		log.Error().Msgf("Error during reading transaction by ID: %d, err: %v", transactionID, err)
		return read.Transactions{}, err
	}

	result := read.Transactions{}
	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.User_id, &result.Description); err != nil {
			log.Error().Msgf("Error during decoding transaction by ID: %d, err: %v", transactionID, err)
			return read.Transactions{}, err
		}
	}
	return read.Transactions{}, nil
}

func (e *DatabaseEngine) UpdateTransaction(ctx context.Context, transactionID int, transaction write.Transactions) error {
	_, err := e.db.ExecContext(
		ctx,
		"UPDATE users SET description = $2, amount = $3, date = $4 WHERE id = $1",
		transactionID, transaction.Description, transaction.Amount, transaction.Date,
	)
	if err != nil {
		log.Error().Msgf("Wrong update: %v", err)
	}
	return err
}

func (e *DatabaseEngine) DeleteTransaction(ctx context.Context, transactionID int) error {
	_, err := e.db.ExecContext(
		ctx,
		"DELETE FROM transactions WHERE id = $1",
		transactionID,
	)
	if err != nil {
		log.Error().Msgf("Wrong delete request: %v", err)
	}
	return err
}

func (e *DatabaseEngine) GetTrasactionByOption(ctx context.Context, field string, value any) (read.Transactions, error) {

	allowed := map[string]bool{
		"id":          true,
		"user_id":     true,
		"description": true,
		"amount":      true,
		"date":        true,
	}

	if !allowed[field] {
		return read.Transactions{}, fmt.Errorf("Ivalid field: %s", field)
	}

	query := fmt.Sprintf(` SELECT id, user_id, description, amount, date FROM transactions WHERE %s = $1`, field)

	rows := e.db.QueryRowContext(ctx, query, value)

	result := read.Transactions{}, nil
	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.User_id, &result.Description, &result.Amount, &result.Date); err != nil {
			log.Error().Msgf("Error during decoding transaction by %s = %v, err: %v")
			return read.Transactions{}, err
		}
		return result, nil
	}
}
