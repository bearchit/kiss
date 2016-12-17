package sql

import "github.com/jmoiron/sqlx"

type Tx struct {
	*sqlx.Tx
}

func (db *DB) Begin() (*Tx, error) {
	tx, err := db.Beginx()
	return &Tx{tx}, err
}

func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}
