package databases

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/caohoangphuctd97/go-test/internal/app/queries"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"

	// postgres driver
	_ "github.com/lib/pq"
)

type (
	Databases struct {
		dig.Out
		Pg *sql.DB `name:"pg"`
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

type Queries struct {
	*queries.BookQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {

	dbConfig := DatabaseCfg{}
	env.Parse(&dbConfig)
	db := openPostgres(&dbConfig)

	return &Queries{
		BookQueries: &queries.BookQueries{DB: db}, // from Book model
	}, nil
}

func openPostgres(p *DatabaseCfg) *sqlx.DB {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.DBUser, p.DBPass, p.Host, p.Port, p.DBName,
	)
	db, err := sqlx.Open("postgres", conn)
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
