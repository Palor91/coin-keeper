package filters

import (
	dbFilters "coin-keeper/internal/database/sql/models/filters"
	httpFilters "coin-keeper/internal/server/http/requests/filters"
	"database/sql"
	"strconv"
	"time"
)

const dateOnlyLayout = "2006-01-02"

func ConvertTransactionFilter(filter httpFilters.TransactionFilter) dbFilters.TransactionFilter {
	dbFilter := dbFilters.TransactionFilter{}

	userID, err := strconv.Atoi(filter.UserID)
	if err != nil {
		userID = 0
	}

	minAmount, err := strconv.ParseFloat(filter.MinAmount, 64)
	if err != nil {
		minAmount = 0
	}

	maxAmount, err := strconv.ParseFloat(filter.MaxAmount, 64)
	if err != nil {
		maxAmount = 0
	}

	dateFrom, err := time.Parse(dateOnlyLayout, filter.DateFrom)
	if err != nil {
		dbFilter.DateFrom = sql.NullTime{Valid: false}
	} else {
		dbFilter.DateFrom = sql.NullTime{Valid: true, Time: dateFrom}
	}

	dateTo, err := time.Parse(dateOnlyLayout, filter.DateTo)
	if err != nil {
		dbFilter.DateTo = sql.NullTime{Valid: false}
	} else {
		dbFilter.DateTo = sql.NullTime{Valid: true, Time: dateTo}
	}

	dbFilter.UserID = userID
	dbFilter.MinAmount = minAmount
	dbFilter.MaxAmount = maxAmount

	return dbFilter
}
