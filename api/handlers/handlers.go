package handlers

import (
	"babylon-stack/api/dao"
	"babylon-stack/api/models"
	"babylon-stack/utilstools"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get ALL Countries
func GetAllCountriesEndPoint(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllCountries()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func GetMinWageEndPoint(w http.ResponseWriter, r *http.Request) {
	payload := utilstools.GetDataXLX()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func GetMinWageEndPointMongo(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetWageMongo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func DeleteWageEndpoint(w http.ResponseWriter, r *http.Request) {
	var wage models.Minimumwage
	_ = json.NewDecoder(r.Body).Decode(&wage)
	dao.DeleteWage(wage)
}

// UpdateWageEndpoint updates a Wage
func UpdateWageEndpoint(w http.ResponseWriter, r *http.Request) {
	wageID := mux.Vars(r)["id"]
	var wage models.Minimumwage
	_ = json.NewDecoder(r.Body).Decode(&wage)
	dao.UpdateWage(wage, wageID)

}
