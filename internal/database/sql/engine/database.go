package engine

import "database/sql"

type DatabaseEngine struct {
	db *sql.DB
}
