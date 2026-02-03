package filters

import "database/sql"

type TransactionFilter struct {
	UserID    int
	MinAmount float64
	MaxAmount float64
	DateFrom  sql.NullTime
	DateTo    sql.NullTime
	Limit     int
	Offset    int
}
