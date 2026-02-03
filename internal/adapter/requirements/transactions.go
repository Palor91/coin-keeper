package requirements

import (
	"coin-keeper/internal/database/sql/models/filters"
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"context"
)

type RequiredTransactions interface {
	CreateTransaction(ctx context.Context, user write.Transactions) error
	ReadTransaction(ctx context.Context, transactionID int) (read.Transactions, error)
	UpdateTransaction(ctx context.Context, transactionID int, transaction write.Transactions) error
	DeleteTransaction(ctx context.Context, transactionID int) error
	GetTrasactionByOption(
		ctx context.Context,
		filter filters.TransactionFilter,
	) ([]read.Transactions, error)
}
