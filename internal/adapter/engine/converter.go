package engine

import (
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"coin-keeper/internal/server/http/requests"
)

func ConvertUserReqToDBW(req requests.User) write.User {
	return write.User{
		Name:     req.Name,
		Login:    req.Login,
		Password: req.Password,
	}
}

func ConvertUserReqToDBR(req requests.User) read.User {
	return read.User{
		ID:    req.ID,
		Name:  req.Name,
		Login: req.Login,
	}
}
