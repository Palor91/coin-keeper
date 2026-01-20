package requests

type Transaction struct {
	ID          int    `json:"id,omitempty"`
	User_id     int    `json:"user_id,omitempty"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
}
