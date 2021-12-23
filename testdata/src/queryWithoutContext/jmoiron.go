package queryWithoutContext

import (
	"context"
)

func warningsJmoiron() {
	var (
		db   SQLDbx
		tx   SQLTxx
		stmt SQLStmt
	)
	tx = db.MustBegin() // want `don't send query to external storage without context`
	db.MustBeginTx(context.Background(), nil)

	stmt = tx.StmtContext(context.Background(), nil)

	{
		db.MustExec(`SELECT 1`)       // want `don't send query to external storage without context`
		db.NamedExec(`SELECT 1`, nil) // want `don't send query to external storage without context`
		db.Exec(`SELECT 1`)           // want `don't send query to external storage without context`

		db.Get(nil, `SELECT 1`)    // want `don't send query to external storage without context`
		db.Select(nil, `SELECT 1`) // want `don't send query to external storage without context`

		db.Prepare(`SELECT 1`)      // want `don't send query to external storage without context`
		db.Preparex(`SELECT 1`)     // want `don't send query to external storage without context`
		db.PrepareNamed(`SELECT 1`) // want `don't send query to external storage without context`

		db.Query(`SELECT 1`)           // want `don't send query to external storage without context`
		db.Queryx(`SELECT 1`, nil)     // want `don't send query to external storage without context`
		db.QueryRow(`SELECT 1`)        // want `don't send query to external storage without context`
		db.QueryRowx(`SELECT 1`, nil)  // want `don't send query to external storage without context`
		db.NamedQuery(`SELECT 1`, nil) // want `don't send query to external storage without context`
	}

	{
		tx.Exec(`SELECT 1`)           // want `don't send query to external storage without context`
		tx.MustExec(`SELECT 1`)       // want `don't send query to external storage without context`
		tx.NamedExec(`SELECT 1`, nil) // want `don't send query to external storage without context`

		tx.Select(nil, `SELECT 1`) // want `don't send query to external storage without context`
		tx.Get(nil, `SELECT 1`)    // want `don't send query to external storage without context`

		tx.Stmtx(nil)     // want `don't send query to external storage without context`
		tx.NamedStmt(nil) // want `don't send query to external storage without context`
		tx.Stmt(nil)      // want `don't send query to external storage without context`

		tx.Prepare(`SELECT 1`)      // want `don't send query to external storage without context`
		tx.Preparex(`SELECT 1`)     // want `don't send query to external storage without context`
		tx.PrepareNamed(`SELECT 1`) // want `don't send query to external storage without context`

		tx.Query(`SELECT 1`)           // want `don't send query to external storage without context`
		tx.Queryx(`SELECT 1`, nil)     // want `don't send query to external storage without context`
		tx.QueryRow(`SELECT 1`)        // want `don't send query to external storage without context`
		tx.QueryRowx(`SELECT 1`, nil)  // want `don't send query to external storage without context`
		tx.NamedQuery(`SELECT 1`, nil) // want `don't send query to external storage without context`
	}

	{
		stmt.Exec()     // want `don't send query to external storage without context`
		stmt.Query()    // want `don't send query to external storage without context`
		stmt.QueryRow() // want `don't send query to external storage without context`
	}

	{
		db.MustExecContext(context.Background(), `SELECT 1`)
		db.NamedExecContext(context.Background(), `SELECT 1`, nil)
		db.ExecContext(context.Background(), `SELECT 1`)

		db.GetContext(context.Background(), nil, `SELECT 1`)
		db.SelectContext(context.Background(), nil, `SELECT 1`)

		db.PrepareContext(context.Background(), `SELECT 1`)
		db.PreparexContext(context.Background(), `SELECT 1`)
		db.PrepareNamedContext(context.Background(), `SELECT 1`)

		db.QueryContext(context.Background(), `SELECT 1`)
		db.QueryxContext(context.Background(), `SELECT 1`, nil)
		db.QueryRowContext(context.Background(), `SELECT 1`)
		db.QueryRowxContext(context.Background(), `SELECT 1`, nil)
		db.NamedQueryContext(context.Background(), `SELECT 1`, nil)
	}

	{
		tx.ExecContext(context.Background(), `SELECT 1`)
		tx.MustExecContext(context.Background(), `SELECT 1`)
		tx.NamedExecContext(context.Background(), `SELECT 1`, nil)

		tx.SelectContext(context.Background(), nil, `SELECT 1`)
		tx.GetContext(context.Background(), nil, `SELECT 1`)

		tx.StmtxContext(context.Background(), nil)
		tx.NamedStmtContext(context.Background(), nil)
		tx.StmtContext(context.Background(), nil)

		tx.PrepareContext(context.Background(), `SELECT 1`)
		tx.PreparexContext(context.Background(), `SELECT 1`)
		tx.PrepareNamedContext(context.Background(), `SELECT 1`)

		tx.QueryContext(context.Background(), `SELECT 1`)
		tx.QueryxContext(context.Background(), `SELECT 1`, nil)
		tx.QueryRowContext(context.Background(), `SELECT 1`)
		tx.QueryRowxContext(context.Background(), `SELECT 1`, nil)
	}

	{
		stmt.QueryRowContext(context.Background())
		stmt.QueryContext(context.Background())
		stmt.ExecContext(context.Background())
	}
}
