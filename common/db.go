package common

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

// GetDb connects to the database and returns the orm.
func GetDb() (*bun.DB, error) {
	dsn, ok := os.LookupEnv("DSN")

	if !ok {
		return nil, fmt.Errorf("DSN not found in env")
	} else {
		db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		return bun.NewDB(db, pgdialect.New()), nil
	}
}
