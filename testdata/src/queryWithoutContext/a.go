package queryWithoutContext

import (
	"database/sql"
)

type decorator struct {
	*sql.DB
}

type decoratorWithParams struct {
	d *sql.DB
}

func warnings() {
	db, _ := sql.Open("PostgreSQL", "test")
	db.Query(`SELECT 1`)    // want `don't send query to external storage without context`
	db.QueryRow(`SELECT 1`) // want `don't send query to external storage without context`
	db.Exec(`SELECT 1`)     // want `don't send query to external storage without context`
	db.Prepare(`SELECT 1`)  // want `don't send query to external storage without context`
	db.Ping()               // want `don't send query to external storage without context`
	db.Begin()              // want `don't send query to external storage without context`

	d := decorator{db}

	d.Query(`SELECT 1`)    // TODO: want `don't send query to external storage without context`
	d.QueryRow(`SELECT 1`) // TODO: want `don't send query to external storage without context`
	d.Exec(`SELECT 1`)     // TODO: want `don't send query to external storage without context`
	d.Prepare(`SELECT 1`)  // TODO: want `don't send query to external storage without context`
	d.Ping()               // TODO: want `don't send query to external storage without context`
	d.Begin()              // TODO: want `don't send query to external storage without context`

	dw := decoratorWithParams{d: db}
	dw.d.Query(`SELECT 1`)    // want `don't send query to external storage without context`
	dw.d.QueryRow(`SELECT 1`) // want `don't send query to external storage without context`
	dw.d.Exec(`SELECT 1`)     // want `don't send query to external storage without context`
	dw.d.Prepare(`SELECT 1`)  // want `don't send query to external storage without context`
	dw.d.Ping()               // want `don't send query to external storage without context`
	dw.d.Begin()              // want `don't send query to external storage without context`
}
