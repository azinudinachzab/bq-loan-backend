package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
