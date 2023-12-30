package database

import (
	"bi/config"
	"context"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Database struct {
	conn              *sql.DB
	B                 sq.StatementBuilderType
	placeholderFormat sq.PlaceholderFormat
}

func New(cfg config.Database) (*Database, error) {
	if !isValidType(cfg.Type) {
		return nil, errors.New(fmt.Sprintf(
			`неизвестный тип базы данных: "%s"`,
			cfg.Type,
		))
	}

	conn, err := sql.Open(cfg.Type, cfg.Dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		conn: conn,
		B: sq.StatementBuilder.
			PlaceholderFormat(getPlaceholderFormat(cfg.Type)),
		placeholderFormat: getPlaceholderFormat(cfg.Type),
	}, nil
}

const (
	Postgres = "postgres"
)

func isValidType(typ string) bool {
	switch typ {
	case Postgres:
		return true
	default:
		return false
	}
}

func getPlaceholderFormat(typ string) sq.PlaceholderFormat {
	switch typ {
	case Postgres:
		return sq.Dollar

	//НЕДОСТУПЕН, ПОСКОЛЬКУ ПРЕДПОЛАГАЕТСЯ,
	//ЧТО ТИП БАЗЫ ДАННЫХ БЫЛ ПРОВЕРЕН ПЕРЕД ВЫЗОВОМ ЭТОЙ ФУНКЦИИ.
	default:
		return nil
	}
}

func (db *Database) PlaceholderFormat() sq.PlaceholderFormat {
	return db.placeholderFormat
}

const (
	txKey = "tx"
)

func (db *Database) Begin(
	ctx context.Context,
	transaction func(context.Context) error,
) error {

	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	if err = transaction(context.WithValue(ctx, txKey, tx)); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (db *Database) Query(
	ctx context.Context,
	query string,
	args ...any,
) (*sql.Rows, error) {

	if ctx.Value(txKey) != nil {
		return ctx.Value(txKey).(*sql.Tx).QueryContext(ctx, query, args...)
	}

	return db.conn.QueryContext(ctx, query, args...)
}

func (db *Database) QueryRow(
	ctx context.Context,
	query string,
	args ...any,
) *sql.Row {

	if ctx.Value(txKey) != nil {
		return ctx.Value(txKey).(*sql.Tx).QueryRowContext(ctx, query, args...)
	}

	return db.conn.QueryRowContext(ctx, query, args...)
}

func (db *Database) Exec(
	ctx context.Context,
	query string,
	args ...any,
) (sql.Result, error) {

	if ctx.Value(txKey) != nil {
		return ctx.Value(txKey).(*sql.Tx).ExecContext(ctx, query, args...)
	}

	return db.conn.ExecContext(ctx, query, args...)
}
