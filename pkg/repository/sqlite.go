package repository

import "github.com/jmoiron/sqlx"

const (
	schedulerTable = "scheduler"
)

type Config struct {
	Driver   string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewSqliteDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, cfg.DBName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
