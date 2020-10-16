package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Data struct {
	ID       string
	Location struct {
		Lat  float32
		Long float32
	}
	DateAdded time.Time
}

type DataRepo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *DataRepo {
	return &DataRepo{DB: db}
}

func (d *DataRepo) Get() []Data {
	var result []Data
	d.DB.Select(&result, "select * from Data")
	return result
}
