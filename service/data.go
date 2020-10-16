package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../repo"
)

type IDataService interface {
	Get()
	Create()
}
type DataService struct {
	DataRepo *repo.DataRepo
}

func NewDataService(datarepo *repo.DataRepo) *DataService {
	return &DataService{DataRepo: datarepo}
}

func (d *DataService) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		result := d.DataRepo.Get()
		responseWithJson(w, 200, result)
		return
	}
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		var data repo.Data
		json.Unmarshal(body, &data)
		err = d.DataRepo.Create(data)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		}
	}
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
