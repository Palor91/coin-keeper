package write

type User struct {
	Name     string `db:"name"`
	Login    string `db:"login"`
	Password string `db:"password"`
}
