package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func main() {
	err := Main()
	if err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	p := os.Getenv("DB_PASSWORD")
	db, err := sqlx.Connect("mysql", fmt.Sprintf("crmuser:%s@(localhost:13307)/crm", p))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(db)
	return nil
}

type member struct {
	ID int `db:"id"`
}
