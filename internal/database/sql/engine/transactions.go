package engine

import (
	"coin-keeper/internal/database/sql/models/filters"
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"context"
	"fmt"
	"strings"

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

func (e *DatabaseEngine) GetTrasactionByOption(
	ctx context.Context,
	filter filters.TransactionFilter,
) ([]read.Transactions, error) {

	filters := make([]string, 0)
	args := make([]any, 0)
	i := 1

	if filter.UserID != 0 {
		filters = append(filters, fmt.Sprintf("user_id = $%d", i))
		args = append(args, filter.UserID)
		i++
	}

	if filter.MinAmount != 0 {
		filters = append(filters, fmt.Sprintf("amount >= $%d", i))
		args = append(args, filter.MinAmount)
		i++
	}

	if filter.MaxAmount != 0 {
		filters = append(filters, fmt.Sprintf("amount <= $%d", i))
		args = append(args, filter.MaxAmount)
		i++
	}

	if filter.DateFrom.Valid {
		filters = append(filters, fmt.Sprintf("date >= $%d", i))
		args = append(args, filter.DateFrom)
		i++
	}

	if filter.DateTo.Valid {
		filters = append(filters, fmt.Sprintf("date <= $%d", i))
		args = append(args, filter.DateTo)
		i++
	}

	// 3. Собираем SQL‑запрос
	query := `
        SELECT id, user_id, description, amount, date,
        FROM transactions
    `
	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)

	// 4. Выполняем запрос
	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error().Msgf("Error during reading transactions: %v", err)
		return nil, err
	}
	defer rows.Close()

	// 5. Сканируем результат
	var result []read.Transactions

	for rows.Next() {
		var t read.Transactions
		if err := rows.Scan(&t.ID, &t.User_id, &t.Description, &t.Amount, &t.Date); err != nil {
			log.Error().Msgf("Error during scanning transaction: %v", err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}
