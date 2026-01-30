package requirements

import (
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/responses"
	"context"
)

type RequiredTrasactions interface {
	CreateTransaction(ctx context.Context, transaction requests.Transaction) error
	ReadTransaction(ctx context.Context, transactionID int) (responses.Transaction, error)
	UpdateTransaction(ctx context.Context, transaction requests.Transaction) error
	DeleteTransaction(ctx context.Context, transactionID int) error
	GetTrasactionByOption(ctx context.Context, transaction requests.Transaction) (responses.Transaction, error)
}
