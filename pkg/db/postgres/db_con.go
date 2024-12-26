package postgres

import (
	"fmt"
	"time"

	"github.com/amankumarsingh77/cmr/config"
	_ "github.com/jackc/pgx/v4/stdlib" // This is important - registers pgx driver
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPsqlDB(c *config.Config) (*sqlx.DB, error) {
	if c.Postgres.PgDriver == "" {
		c.Postgres.PgDriver = "pgx"
	}

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=require password=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Name,
		c.Postgres.Password,
	)

	db, err := sqlx.Connect(c.Postgres.PgDriver, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return db, nil
}