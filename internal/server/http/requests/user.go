package requests

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
