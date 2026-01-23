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
