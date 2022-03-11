package db

import (
	"database/sql"
	"store/support"

	_ "github.com/lib/pq"
)

type PostgresDb struct {
	db *sql.DB
}

func connect() *sql.DB {
	config := support.LoadConfig()
	db, err := sql.Open("postgres", config.DB)
	support.PanicError(err)

	return db
}

func close(db *sql.DB) {
	defer db.Close()
}

func ExecPreparedQuery(stmt string, args ...interface{}) {
	db := connect()
	query, queryErr := db.Prepare(stmt)
	support.PanicError(queryErr)
	_, execErr := query.Exec(args...)
	support.PanicError(execErr)
	close(db)
}

func ExecQuery(query string, args ...interface{}) *sql.Rows {
	db := connect()
	results, err := db.Query(query, args...)
	support.PanicError(err)
	close(db)
	return results
}
