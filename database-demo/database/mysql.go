package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(dbURL string) (*DBManager, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &DBManager{
		db: db,
	}, nil
}

func (m *DBManager) Close() {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			log.Printf("failed to close database connection: %v", err)
		}
	}
}

func (m *DBManager) Execute(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %v", err)
	}

	return result, nil
}

func (m *DBManager) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.db.QueryRow(query, args...)
}

func (m *DBManager) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.db.Query(query, args...)
}
