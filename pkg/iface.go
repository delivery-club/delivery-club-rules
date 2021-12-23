package pkg

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SQLDb interface {
	SQLPing
	SQLQuery
	SQLQueryRow
	SQLExec
	SQLBeginTx
	SQLPrepare
}

type SQLTx interface {
	SQLQuery
	SQLExec
	SQLQueryRow
	SQLPrepare
	SQLCreateStmt
}

type SQLStmt interface {
	SQLStmtQuery
	SQLStmtQueryRow
	SQLStmtExec
}

type SQLDbx interface {
	SQLPing
	SQLQueryx
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	SQLQueryRow
	SQLExecx
	SQLSelectx
	SQLPreparex
	SQLBeginTxx
}

type SQLTxx interface {
	SQLTx
	SQLExecx
	SQLSelectx
	SQLPreparex
	SQLQueryx
	SQLCreateStmtx
}

type SQLPing interface {
	Ping() error
	PingContext(ctx context.Context) error
}

type SQLCreateStmt interface {
	Stmt(stmt *sql.Stmt) *sql.Stmt
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}

type SQLQuery interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type SQLStmtQuery interface {
	QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
	Query(args ...interface{}) (*sql.Rows, error)
}

type SQLQueryRow interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type SQLStmtQueryRow interface {
	QueryRow(args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
}

type SQLBeginTx interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
}

type SQLPrepare interface {
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type SQLExec interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type SQLStmtExec interface {
	Exec(args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
}

type SQLExecx interface {
	SQLExec
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type SQLSelectx interface {
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(context context.Context, dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(context context.Context, dest interface{}, query string, args ...interface{}) error
}

type SQLPreparex interface {
	sqlx.Preparer
	sqlx.PreparerContext
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(context context.Context, query string) (*sqlx.NamedStmt, error)
}

type SQLBeginTxx interface {
	MustBegin() *sqlx.Tx
	MustBeginTx(ctx context.Context, opts *sql.TxOptions) *sqlx.Tx
	SQLBeginTx
}

type SQLQueryx interface {
	sqlx.Queryer
	sqlx.QueryerContext
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}

type SQLCreateStmtx interface {
	SQLCreateStmt
	Stmtx(stmt interface{}) *sqlx.Stmt
	StmtxContext(ctx context.Context, stmt interface{}) *sqlx.Stmt
	NamedStmt(stmt *sqlx.NamedStmt) *sqlx.NamedStmt
	NamedStmtContext(ctx context.Context, stmt *sqlx.NamedStmt) *sqlx.NamedStmt
}
