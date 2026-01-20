package engine

import (
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/responses"
	"context"
)

func (a *Adapter) CreateTransaction(ctx context.Context, req requests.Transaction) error {
	transactionDB := ConvertTransactionReqToDBWrite(req)

	return a.dbEngine.CreateTransaction(ctx, transactionDB)
}

func (a *Adapter) ReadTransaction(ctx context.Context, id int) (responses.Transaction, error) {
	transaction, err := a.dbEngine.ReadTransaction(ctx, id)
	if err != nil {
		return responses.Transaction{}, err
	}
	return ConvertTransactionReqDBToResp(transaction), nil
}

func (a *Adapter) UpdateTransaction(ctx context.Context, req requests.Transaction) error {
	transactionDB := ConvertTransactionReqToDBWrite(req)

	return a.dbEngine.UpdateTransaction(ctx, req.ID, transactionDB)
}

func (a *Adapter) DeleteTransaction(ctx context.Context, id int) error {

	return a.dbEngine.DeleteTransaction(ctx, id)
}
