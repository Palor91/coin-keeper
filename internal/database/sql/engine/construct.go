package engine

import "database/sql"

func NewDatabaseEngine(fileName string) (*DatabaseEngine, error) {
	db, err := sql.Open("sqlite3", fileName)

	if err != nil {
		return nil, err
	}
	return &DatabaseEngine{
		db: db,
	}, nil
}
