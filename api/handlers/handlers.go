package handlers

import (
	"babylon-stack/api/dao"
	"babylon-stack/api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-resty/resty/v2"
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

func GetCurrency(w http.ResponseWriter, req *http.Request) {
	//var currency string = "USD_XCD"

	vars := mux.Vars(req)
	country1 := vars["item1"]
	country2 := vars["item2"]

	fmt.Println("Vars ", vars)
	fmt.Println("Country 1", country1)
	fmt.Println("Country 2", country2)

	var country models.Country

	payload := dao.GetAll(country)

	fmt.Println("Tomaaa", payload)
	fmt.Println("Type Recived", reflect.TypeOf(payload))
	/*for _, p := range payload {
		if p.ID == country1 && p.ID == country2 {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Person not found")*/

	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		//Get("https://free.currconv.com/api/v7/convert?compact=ultra&apiKey=341365d39b96b88174d7&q=" + currency)
		Get("http://localhost:8020/country/5d7f6b2b57d5104f58e53d2a")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
	fmt.Println()

}
