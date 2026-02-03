package engine

import (
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/responses"
	"time"
)

var layout = "2006-01-02 15:04:05"

func ConvertTransactionReqToDBWrite(req requests.Transaction) write.Transactions {
	t, err := time.Parse(layout, req.Date)
	if err != nil {
		panic(err)
	}

	return write.Transactions{
		ID:          req.ID,
		User_id:     req.User_id,
		Description: req.Description,
		Amount:      req.Amount,
		Date:        t,
	}
}

func ConvertTransactionReqDBToResp(u read.Transactions) responses.Transaction {
	return responses.Transaction{
		ID:          u.ID,
		User_id:     u.User_id,
		Description: u.Description,
		Amount:      u.Amount,
		Date:        u.Date.Format(layout),
	}
}

func ConvertTransactionsDBtoResps(model ...read.Transactions) []responses.Transaction {
	var resp = make([]responses.Transaction, len(model))
	for i, u := range model {
		resp[i] = ConvertTransactionReqDBToResp(u)
	}
	return resp
}
