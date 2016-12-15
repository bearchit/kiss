package sql

import "github.com/jmoiron/sqlx"

type arg map[string]interface{}

func (db *DB) GetRowsNamedIn(query string, arg arg) (rows *sqlx.Rows, err error) {
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return
	}

	query = db.Rebind(query)
	rows, err = db.Queryx(query, args...)

	return
}

func (db *DB) GetRowNamedIn(query string, arg arg) (*sqlx.Row, error) {
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	return db.QueryRowx(query, args...), nil
}

func (db *DB) getRowNamed(query string, arg arg) (*sqlx.Row, error) {
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return nil, err
	}

	return db.QueryRowx(query, args...), nil
}
