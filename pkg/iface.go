package pkg

import (
	"context"
	"database/sql"
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
