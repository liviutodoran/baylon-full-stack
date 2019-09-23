package handlers

import (
	"encoding/json"
	"net/http"
	"babylon-stack/api/models"
	"babylon-stack/api/dao"
)

var countries []models.Country

// Get ALL Countries
func GetAllCountriesEndPoint(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllCountries()
	json.NewEncoder(w).Encode(payload)
}