package handlers

import (
	"babylon-stack/api/dao"
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

		payload := dao.UpdateItem(elem, itemID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	})
}

func AddItem(data interface{}) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		types := reflect.TypeOf(data)
		elem := reflect.New(types).Interface()

		_ = json.NewDecoder(req.Body).Decode(&elem)
		dao.AddItem(elem)
		json.NewEncoder(w).Encode(elem)
	})

}

func DeleteItem(data interface{}) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		types := reflect.TypeOf(data)
		elem := reflect.New(types).Interface()

		_ = json.NewDecoder(req.Body).Decode(&elem)
		dao.DeleteItem(elem)
	})

}
