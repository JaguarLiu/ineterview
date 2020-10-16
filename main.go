package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Connect("mysql", "test:test@(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	svr := service.NewDataService(repo.New(db))
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8000s", nil)
}
