package postgres

import (
	"fmt"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB initializes the database connection
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(pg.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// NewDB initializes the database connection
func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Errorf("failed to create sqlmock: %w", err))
	}
	db, err := gorm.Open(pg.New(pg.Config{Conn: sqlDB}))
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	return db, mock
}
