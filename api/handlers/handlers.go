package handlers

import (
	"babylon-stack/api/dao"
	"babylon-stack/api/models"
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

func GetCountryEndPoint(w http.ResponseWriter, r *http.Request) {
	countryID := mux.Vars(r)["id"]
	var country models.Country
	payload := dao.GetCountry(country, countryID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

// UpdateWageEndpoint updates a Wage
func UpdateCountryEndpoint(w http.ResponseWriter, r *http.Request) {
	countryID := mux.Vars(r)["id"]
	var country models.Country
	_ = json.NewDecoder(r.Body).Decode(&country)
	payload := dao.UpdateCountry(country, countryID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)

}

func GetAllMinWageEndPoint(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllWage()
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
	payload := dao.UpdateWage(wage, wageID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)

}
