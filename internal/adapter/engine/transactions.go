package engine

import (
	"context"

	adapterConverters "coin-keeper/internal/converters"
	filtersConverters "coin-keeper/internal/converters/filters"
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/requests/filters"
	"coin-keeper/internal/server/http/responses"
)

func (a *Adapter) CreateTransaction(ctx context.Context, req requests.Transaction) error {
	transactionDB := adapterConverters.ConvertTransactionReqToDBWrite(req)

	return a.dbEngine.CreateTransaction(ctx, transactionDB)
}

func (a *Adapter) ReadTransaction(ctx context.Context, id int) (responses.Transaction, error) {
	transaction, err := a.dbEngine.ReadTransaction(ctx, id)
	if err != nil {
		return responses.Transaction{}, err
	}
	return adapterConverters.ConvertTransactionReqDBToResp(transaction), nil
}

func (a *Adapter) UpdateTransaction(ctx context.Context, req requests.Transaction) error {
	transactionDB := adapterConverters.ConvertTransactionReqToDBWrite(req)

	return a.dbEngine.UpdateTransaction(ctx, req.ID, transactionDB)
}

func (a *Adapter) DeleteTransaction(ctx context.Context, id int) error {
	return a.dbEngine.DeleteTransaction(ctx, id)
}

func (a *Adapter) GetTransactionByOption(
	ctx context.Context,
	filter filters.TransactionFilter,
) ([]responses.Transaction, error) {
	dbFilter := filtersConverters.ConvertTransactionFilter(filter)
	result, err := a.dbEngine.GetTrasactionByOption(ctx, dbFilter)
	if err != nil {
		return nil, err
	}
	return adapterConverters.ConvertTransactionsDBtoResps(result...), nil
}
