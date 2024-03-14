package databases

import (
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/dig"

	// postgres driver
	_ "github.com/lib/pq"
)

type (
	Databases struct {
		dig.Out
		Pg *sql.DB `name:"pg"`
	}

	// Databases setup output
	DatabaseCfgs struct {
		dig.In
		Pg *DatabaseCfg `name:"pg"`
	}
	// DatabaseCfg is MySQL configuration
	// @envconfig (prefix:"PG" ctor:"pg")
	// // @envconfig (prefix:"MYSQL" ctor:"mysql")
	DatabaseCfg struct {
		DBName string `envconfig:"DBNAME" required:"true" default:"dbname"`
		DBUser string `envconfig:"DBUSER" required:"true" default:"dbuser"`
		DBPass string `envconfig:"DBPASS" required:"true" default:"dbpass"`
		Host   string `envconfig:"HOST" required:"true" default:"localhost"`
		Port   string `envconfig:"PORT" required:"true" default:"9999"`

		MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"30" required:"true"`
		MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"6" required:"true"`
		ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"30m" required:"true"`
	}
)

// NewDatabases return new instance of databases
// @ctor
func NewDatabases(cfgs DatabaseCfgs) Databases {
	return Databases{
		Pg: openPostgres(cfgs.Pg),
	}
}

func openPostgres(p *DatabaseCfg) *sql.DB {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.DBUser, p.DBPass, p.Host, p.Port, p.DBName,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Errorf("postgres: %s", err.Error())
	}

	db.SetConnMaxLifetime(p.ConnMaxLifetime)
	db.SetMaxIdleConns(p.MaxIdleConns)
	db.SetMaxOpenConns(p.MaxOpenConns)

	if err = db.Ping(); err != nil {
		fmt.Errorf("postgres: %s", err.Error())
	}

	return db
}
