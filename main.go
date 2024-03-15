package main

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

type config struct {
	Home   string `env:"HOME"`
	DBName string `env:"DBNAME" envDefault:"dbname"`
	DBUser string `env:"DBUSER" envDefault:"dbuser"`
	DBPass string `env:"DBPASS" envDefault:"dbpass"`
	Host   string `env:"HOST" envDefault:"localhost"`
	Port   string `env:"PORT" envDefault:"9999"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}
