package requirements

import (
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/responses"
	"context"
)

type RequiredUsers interface {
	CreateUser(ctx context.Context, user requests.User) error
	ReadUser(ctx context.Context, userID int) (responses.User, error)
	UpdateUser(ctx context.Context, user requests.User) error
	DeleteUser(ctx context.Context, userID int) error
}
