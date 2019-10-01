package handlers

import (
	"babylon-stack/api/dao"
	"babylon-stack/api/models"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

// Get ALL Items
func GetAll(data interface{}) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		payload := dao.GetAll(data)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	})
}

func GetItem(data interface{}) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		itemID := mux.Vars(req)["id"]
		payload := dao.GetItem(data, itemID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	})
}

func UpdateItem(data interface{}) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		itemID := mux.Vars(req)["id"]

		types := reflect.TypeOf(data)
		elem := reflect.New(types).Interface()

		_ = json.NewDecoder(req.Body).Decode(elem)
		payload := dao.UpdateItem(data, itemID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	})
}

func AddCountryEndpoint(w http.ResponseWriter, r *http.Request) {
	var country models.Country
	_ = json.NewDecoder(r.Body).Decode(&country)
	dao.AddCountry(country)
	json.NewEncoder(w).Encode(country)
}

func DeleteCountryEndpoint(w http.ResponseWriter, r *http.Request) {
	var country models.Country
	_ = json.NewDecoder(r.Body).Decode(&country)
	dao.DeleteCountry(country)
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

func AddMinWageEndpoint(w http.ResponseWriter, r *http.Request) {
	var wage models.Minimumwage
	_ = json.NewDecoder(r.Body).Decode(&wage)
	dao.AddWage(wage)
	json.NewEncoder(w).Encode(wage)
}
