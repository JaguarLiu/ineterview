package service

import (
	"/interview/repo"
	"net/http"
)

type IDataService interface {
	Get() []repo.Data
	Create() error
}
type DataService struct {
	DataRepo repo.DataRepo
}

func NewDataService(datarepo *repo.DataRepo) *DataService {
	return &DataService{DataRepo: datarepo}
}

func (d *DataService) Get(w http.ResponseWriter, r *http.Request) []repo.Data {
	result := d.DataRepo.Get()
	return result
}
