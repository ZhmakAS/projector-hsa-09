package main

import (
	"github.com/caarlos0/env/v6"
)

type Env struct {
	MySQLURL      string `env:"MYSQL_URL,required"`
	BatchCount    int64  `env:"BATCH_COUNT,required"`
	RecordsNumber int64  `env:"RECORDS_NUMBER,required"`
}

func (e *Env) Parse() error {
	if err := env.Parse(e); err != nil {
		return err
	}
	return nil
}
