package db

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gitlab.devgroup.tech/shkolkovo/romanych/helpers"
	"time"
)

var db *sqlx.DB

func ConnectDatabase() (err error) {
	dataSourceName := helpers.GetEnvDefault("DATABASE", "")
	if dataSourceName == "" {
		return errors.New("no DATABASE env set")
	}
	if db, err = sqlx.Connect("mysql", dataSourceName); err != nil {
		return err
	}

	// Set the maximum number of concurrently open connections to 50. Setting this
	// to less than or equal to 0 will mean there is no maximum limit (which
	// is also the default setting).
	db.SetMaxOpenConns(helpers.GetEnvAsInt("DB_MAX_OPEN_CONNECTIONS", 50))

	// Set the maximum number of concurrently idle connections to 50. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	db.SetMaxIdleConns(helpers.GetEnvAsInt("DB_MAX_IDLE_CONNECTIONS", 50))

	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	db.SetConnMaxLifetime(time.Minute * time.Duration(helpers.GetEnvAsInt("DB_CONNECTION_MAX_LIFETIME", 10)))

	db.MapperFunc(func(column string) string {
		return column
	})

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}


func Database() *sqlx.DB {
	return db
}