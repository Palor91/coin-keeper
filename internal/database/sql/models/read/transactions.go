package read

type Transactions struct {
	ID          int    `db:"id"`
	User_id     int    `db:"user_id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
}
