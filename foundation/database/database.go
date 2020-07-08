// Package database provides support for access the database.
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq" // The database driver in use.
	// "go.opentelemetry.io/otel/api/global"
)

// Config is the required properties to use the database.
type Config struct {
	User       string
	Password   string
	Host       string
	Name       string
	Port       string
	DisableTLS bool
}

// Open knows how to open a database connection based on the configuration.
func Open(cfg Config) (*sql.DB, error) {

	connstring := cfg.buildConnectionString()

	db, err := sql.Open("mysql", connstring)
	if err != nil {
		return nil, err
	}
	if err := StatusCheck(db); err != nil {
		return nil, err
	}
	return sql.Open("mysql", connstring)
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(db *sql.DB) error {
	// ctx, span := global.Tracer("service").Start(ctx, "platform.database.statuscheck")
	// defer span.End()

	// Run a simple query to determine connectivity. The db has a "Ping" method
	// but it can false-positive when it was previously able to talk to the
	// database but the database has since gone away. Running this query forces a
	// round trip to the database.
	const q = `SELECT true`
	var tmp bool
	// return db.QueryRow(q).Scan(&tmp)
	err := db.QueryRow(q).Scan(&tmp)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) buildConnectionString() string {
	return fmt.Sprintf(cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Name)
}
