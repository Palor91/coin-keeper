package engine

import (
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"coin-keeper/internal/server/http/requests"
	"context"
)

func (a *Adapter) CreateUser(ctx context.Context, req requests.User) error {
	userDB := ConvertUserReqToDBW(req)

	return a.dbEngine.CreateUser(ctx, userDB)
}

func (a *Adapter) ReadUser(ctx context.Context, req requests.User) (read.User, error) {
	userDB := ConvertUserReqToDBR(req)

	return a.dbEngine.ReadUser(ctx, userDB.ID)
}

func (a *Adapter) UpdateUser(ctx context.Context, req requests.User) error {
	userDB := ConvertUserReqToDBR(req)

	return a.dbEngine.UpdateUser(ctx, userDB.ID, write.User{})
}

func (a *Adapter) DeleteUser(ctx context.Context, req requests.User) error {
	userDB := ConvertUserReqToDBR(req)

	return a.dbEngine.DeleteUser(ctx, userDB.ID)
}
