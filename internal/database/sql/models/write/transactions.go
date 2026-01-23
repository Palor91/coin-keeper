package write

import "time"

type Transactions struct {
	ID          int       `db:"id"`
	User_id     int       `db:"user_id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
}
