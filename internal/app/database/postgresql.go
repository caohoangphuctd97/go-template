package databases

import (
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/dig"

	// postgres driver

	"github.com/caarlos0/env/v10"
	_ "github.com/lib/pq"
)

type (
	// Databases setup output
	Databases struct {
		dig.Out
		Pg *sql.DB `name:"pg"`
		// MySQL *sql.DB `name:"mysql"`
	}
	DatabaseCfgs struct {
		dig.In
		Pg *DatabaseCfg `name:"pg"`
		// Mysql *DatabaseCfg `name:"mysql"`
	}

	DatabaseCfg struct {
		DBName string `env:"DBNAME" envDefault:"dbname"`
		DBUser string `env:"DBUSER" envDefault:"dbuser"`
		DBPass string `env:"DBPASS" envDefault:"dbpass"`
		Host   string `env:"HOST" envDefault:"localhost"`
		Port   string `env:"PORT" envDefault:"9999"`

		MaxOpenConns    int           `env:"MAX_OPEN_CONNS" envDefault:"30"`
		MaxIdleConns    int           `env:"MAX_IDLE_CONNS" envDefault:"6"`
		ConnMaxLifetime time.Duration `env:"CONN_MAX_LIFETIME" envDefault:"30m"`
	}
)

func NewDatabases() Databases {
	cfg := DatabaseCfg{}
	env.Parse(&cfg)
	return Databases{
		Pg: openPostgres(&cfg),
		// MySQL: openMySQL(cfgs.Mysql),
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
