package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/tevjef/go-runtime-metrics/expvar"
)

func main() {
	var cfg Env
	if err := cfg.Parse(); err != nil {
		panic(err)
	}

	faker := gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(faker)

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		osCall := <-c
		log.Printf("Stop system call:%+v", osCall)
		cancel()
	}()

	db, err := initMySQL(ctx, cfg.MySQLURL)
	if err != nil {
		log.Println("Failed to connect to MySQL")
		panic(err)
	}

	var count int64 = 0
	for {
		if count >= cfg.RecordsNumber {
			break
		}

		users := make([]User, 0, cfg.BatchCount)
		for i := 0; i < int(cfg.BatchCount); i++ {
			users = append(users, NewUser())
		}
		if err := saveBatchUsers(db, users); err != nil {
			panic(err)
		}

		fmt.Printf("Inserted %d records\n", count)
		count += cfg.BatchCount
	}
	fmt.Println("Successfully generate data!")
}

func saveBatchUsers(db *sqlx.DB, users []User) error {
	columns := []string{
		"first_name",
		"last_name",
		"phone",
		"birth_date",
	}

	insertQuery := `INSERT IGNORE INTO users (` + strings.Join(columns, ",") + `) 
		VALUES (:` + strings.Join(columns, ",:") + `)`

	_, err := db.NamedExec(insertQuery, users)
	if err != nil {
		return err
	}

	return nil
}

func initMySQL(ctx context.Context, uri string) (*sqlx.DB, error) {
	connect, err := sqlx.Open("mysql", uri)
	if err != nil {
		log.Printf("Failed to open mysql connection: %s", err)
		return nil, err
	}

	if err := connect.Ping(); err != nil {
		log.Printf("Could not ping mysql: %s", err)
		return nil, err
	}

	return connect, nil
}
