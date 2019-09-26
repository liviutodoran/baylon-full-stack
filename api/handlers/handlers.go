package handlers

import (
	"babylon-stack/api/dao"
	"babylon-stack/api/models"
	"babylon-stack/utilstools"
	"encoding/json"
	"net/http"
)

var countries []models.Country

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

