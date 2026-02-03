package filters

import "net/http"

type TransactionFilter struct {
	UserID    string
	MinAmount string
	MaxAmount string
	DateFrom  string
	DateTo    string
	Limit     string
	Offset    string
}

func NewTransactionFilter(req *http.Request) TransactionFilter {
	return TransactionFilter{
		UserID:    req.URL.Query().Get("user_id"),
		MinAmount: req.URL.Query().Get("min_amount"),
		MaxAmount: req.URL.Query().Get("max_amount"),
		DateFrom:  req.URL.Query().Get("date_from"),
		DateTo:    req.URL.Query().Get("date_to"),
		Limit:     req.URL.Query().Get("limit"),
		Offset:    req.URL.Query().Get("offset"),
	}
}
