package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/bearchit/kiss/log"
)

type DB struct {
	*sqlx.DB

	Logger *log.Logger
}

type Config struct {
	Host         string
	Port         uint
	Name         string
	User         string
	Password     string
	Charset      string
	Location     string
	MaxIdleConns int
	MaxOpenConns int
	MapperFunc   func(string) string
}

func OpenMySQL(c *Config) (*DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.Charset,
		c.Location,
	))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	if c.MapperFunc != nil {
		db.MapperFunc(c.MapperFunc)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
