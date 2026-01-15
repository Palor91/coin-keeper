package read

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Login string `db:"login"`
}
