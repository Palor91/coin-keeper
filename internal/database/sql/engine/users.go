package engine

import (
	"coin-keeper/internal/database/sql/models/read"
	"coin-keeper/internal/database/sql/models/write"
	"context"

	"github.com/rs/zerolog/log"
)

func (e *DatabaseEngine) CreateUser(ctx context.Context, user write.User) error {
	_, err := e.db.ExecContext(
		ctx,
		"INSERT INTO users (name, login, password) VALUES ($1, $2, $3)",
		user.Name, user.Login, user.Password,
	)
	if err != nil {
		log.Error().Msgf("Wrong creation: %v", err)
	}
	return err
}

func (e *DatabaseEngine) ReadUser(ctx context.Context, userID int) (read.User, error) {
	rows, err := e.db.QueryContext(
		ctx,
		"SELECT (id, name, login) FROM users WHERE id = $1",
		userID,
	)
	if err != nil {
		log.Error().Msgf("Error during reading user by ID: %d, err: %v", userID, err)
		return read.User{}, err
	}

	result := read.User{}
	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.Name, &result.Login); err != nil {
			log.Error().Msgf("Error during decoding user by ID: %d, err: %v", userID, err)
			return read.User{}, err
		}
	}
	return result, nil
}

func (e *DatabaseEngine) UpdateUser(ctx context.Context, userID int, user write.User) error {
	_, err := e.db.ExecContext(
		ctx,
		"UPDATE users SET name = $2, login = $3, password = $4 WHERE id = $1",
		userID, user.Name, user.Login, user.Password,
	)
	if err != nil {
		log.Error().Msgf("Wrong update: %v", err)
	}
	return err
}

func (e *DatabaseEngine) DeleteUser(ctx context.Context, userID int) error {
	_, err := e.db.ExecContext(
		ctx,
		"DELETE FROM users WHERE id = $1",
		userID,
	)
	if err != nil {
		log.Error().Msgf("Wrong delete request: %v", err)
	}
	return err
}
