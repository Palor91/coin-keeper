package requirements

import (
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"context"
)

type RequiredUsers interface {
	CreateUser(ctx context.Context, user write.User) error
	ReadUser(ctx context.Context, userID int) (read.User, error)
	UpdateUser(ctx context.Context, userID int, user write.User) error
	DeleteUser(ctx context.Context, userID int) error
}
