package engine

import (
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/responses"
)

func ConvertUserReqToDBWrite(req requests.User) write.User {
	return write.User{
		Name:     req.Name,
		Login:    req.Login,
		Password: req.Password,
	}
}

func ConvertUserDBToResp(u read.User) responses.User {
	return responses.User{
		ID:    u.ID,
		Name:  u.Name,
		Login: u.Login,
	}
}

func ConvertTransactionReqToDBWrite(req requests.Transaction) write.Transactions {
	return write.Transactions{
		ID:          req.ID,
		User_id:     req.User_id,
		Description: req.Description,
		Amount:      req.Amount,
		Date:        req.Date,
	}
}

func ConvertTransactionReqDBToResp(u read.Transactions) responses.Transaction {
	return responses.Transaction{
		ID:          u.ID,
		User_id:     u.User_id,
		Description: u.Description,
		Amount:      u.Amount,
		Date:        u.Date,
	}
}
