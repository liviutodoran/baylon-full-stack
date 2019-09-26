package main

import (
	"babylon-stack/api/handlers"
	"babylon-stack/utilstools"
	"fmt"
	"log"
	"net/http"
    "encoding/json"
	"github.com/gorilla/mux"	
)

func main() {

	/*excelFileName := "./stuff/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)

	var elements []models.MinWage

	if err != nil {
		log.Fatal(err)
	}
	for _, sheet := range xlFile.Sheets {

		for _, row := range sheet.Rows[1:] {

			elements = append(elements, models.MinWage{
				Country:      row.Cells[0].String(),
				Year:         row.Cells[1].String(),
				Local_Amount: row.Cells[2].String(),
				USD:          row.Cells[3].String(),
			})
			fmt.Println("\n", row.Cells[0])
			fmt.Println("\n", row.Cells[1])
			fmt.Println("\n", row.Cells[2])
			fmt.Println("\n", row.Cells[3])
			/*for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}*/
	elements := utilstools.GetDataXLX	
	s, _ := json.MarshalIndent(elements, "", "\t")
	 fmt.Printf(string(s))



	router := mux.NewRouter()
	router.HandleFunc("/countries", handlers.GetAllCountriesEndPoint).Methods("GET")
	router.HandleFunc("/wage", handlers.GetMinWageEndPoint).Methods("GET")
	//ilie := utilstools.getDataXLX;
	fmt.Println("Starting server on port 8020...")
	log.Fatal(http.ListenAndServe(":8020", router))
}
