package migrate

import (
	"database/sql"
	"fmt"
	"os"
	"sort"
)

type Migration struct {
	Version  int
	Name     string
	UpFile   string
	DownFile string
}

type Runner struct {
	migrations []Migration
}

func NewRunner(migs []Migration) *Runner {
	sort.Slice(migs, func(i, j int) bool {
		return migs[i].Version < migs[j].Version
	})
	return &Runner{migrations: migs}
}

func (r *Runner) ensureVersionTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS schema_version (
            version INT NOT NULL
        );
    `)
	if err != nil {
		return err
	}

	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM schema_version;`).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = db.Exec(`INSERT INTO schema_version (version) VALUES (0);`)
	}

	return err
}

func (r *Runner) CurrentVersion(db *sql.DB) (int, error) {
	var v int
	err := db.QueryRow(`SELECT version FROM schema_version LIMIT 1;`).Scan(&v)
	return v, err
}

func (r *Runner) SetVersion(db *sql.DB, v int) error {
	_, err := db.Exec(`UPDATE schema_version SET version = $1;`, v)
	return err
}

func readSQL(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r *Runner) Up(db *sql.DB) error {
	if err := r.ensureVersionTable(db); err != nil {
		return err
	}

	current, err := r.CurrentVersion(db)
	if err != nil {
		return err
	}

	for _, m := range r.migrations {
		if m.Version == current+1 {
			sqlText, err := readSQL(m.UpFile)
			if err != nil {
				return err
			}

			if _, err := db.Exec(sqlText); err != nil {
				return err
			}

			return r.SetVersion(db, m.Version)
		}
	}

	fmt.Println("No more migrations to apply")
	return nil
}

func (r *Runner) Down(db *sql.DB) error {
	if err := r.ensureVersionTable(db); err != nil {
		return err
	}

	current, err := r.CurrentVersion(db)
	if err != nil {
		return err
	}

	for i := len(r.migrations) - 1; i >= 0; i-- {
		m := r.migrations[i]
		if m.Version == current {
			sqlText, err := readSQL(m.DownFile)
			if err != nil {
				return err
			}

			if _, err := db.Exec(sqlText); err != nil {
				return err
			}

			return r.SetVersion(db, current-1)
		}
	}

	fmt.Println("Already at version 0")
	return nil
}
