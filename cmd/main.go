package main

import (
	"log"
	"net/http"

	"../repo"
	"../service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Connect("mysql", "root@(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	svr := service.NewDataService(repo.New(db))
	http.HandleFunc("/data", svr.Handle)
	http.ListenAndServe(":8000", nil)
}
