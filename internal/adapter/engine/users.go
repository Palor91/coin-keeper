package engine

import (
	engine "coin-keeper/internal/converters"
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/responses"
	"context"
)

func (a *Adapter) CreateUser(ctx context.Context, req requests.User) error {
	userDB := engine.ConvertUserReqToDBWrite(req)

	return a.dbEngine.CreateUser(ctx, userDB)
}

func (a *Adapter) ReadUser(ctx context.Context, id int) (responses.User, error) {
	user, err := a.dbEngine.ReadUser(ctx, id)
	if err != nil {
		return responses.User{}, err
	}
	return engine.ConvertUserDBToResp(user), nil
}

func (a *Adapter) UpdateUser(ctx context.Context, req requests.User) error {
	userDB := engine.ConvertUserReqToDBWrite(req)

	return a.dbEngine.UpdateUser(ctx, req.ID, userDB)
}

func (a *Adapter) DeleteUser(ctx context.Context, id int) error {

	return a.dbEngine.DeleteUser(ctx, id)
}
