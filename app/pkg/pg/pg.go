package pg

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ping (db *sql.DB) (err error) {
	err = db.Ping()
	return err
}