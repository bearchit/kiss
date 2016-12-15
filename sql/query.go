package sql

import (
	"fmt"

	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/Sirupsen/logrus"
)

type queryFn func(interface{}, string, ...interface{}) error

func (db *DB) query(v interface{}, q squirrel.SelectBuilder, fn queryFn) error {
	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	db.Logger.WithFields(logrus.Fields{
		"query": query,
		"args":  fmt.Sprint(args...),
	}).Debug("DB")

	return fn(v, query, args...)
}

func (db *DB) Get(v interface{}, q squirrel.SelectBuilder) error {
	return db.query(v, q, db.DB.Get)
}

func (db *DB) Select(v interface{}, q squirrel.SelectBuilder) error {
	return db.query(v, q, db.DB.Select)
}

func (db *DB) Update(q squirrel.UpdateBuilder) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	db.Logger.WithFields(logrus.Fields{
		"query": query,
		"args":  fmt.Sprint(args...),
	}).Debug("DB")

	return db.Exec(query, args...)
}
