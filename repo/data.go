package repo

import (
	"github.com/jmoiron/sqlx"
)

type Location struct {
	Lat  float32
	Long float32
}
type RespData struct {
	ID        string
	Location  Location
	DateAdded string
}
type Data struct {
	ID        string  `db:"ID"`
	Lat       float32 `db:"Lat"`
	Long      float32 `db:"Long"`
	DateAdded string  `db:"DateAdded"`
}
type DataRepo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *DataRepo {
	return &DataRepo{DB: db}
}

func (d *DataRepo) Get() []RespData {
	fetch := []Data{}
	err := d.DB.Select(&fetch, "SELECT * FROM Data")
	if err != nil {
		return nil
	}
	var result []RespData
	for _, f := range fetch {
		result = append(result, RespData{ID: f.ID, Location: Location{Lat: f.Lat, Long: f.Long}, DateAdded: f.DateAdded})
	}
	return result
}
func (d *DataRepo) Create(data Data) error {
	sqlStr := "INSERT INTO `test`.`Data` (`ID`,`Lat`,`Long`) VALUES (?,?,?);"
	_, err := d.DB.Exec(sqlStr, data.ID, data.Lat, data.Long)
	if err != nil {
		return err
	}
	return nil
}
