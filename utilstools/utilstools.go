package utilstools

import (
	"babylon-stack/api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tealeg/xlsx"
)

func GetDataXLX() []models.Minimumwage {
	excelFileName := "./stuff/National-Minimum-Wage.xlsx"
	var elements []models.Minimumwage
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal("Aqui", err)
	}
	for _, sheet := range xlFile.Sheets {

		for _, row := range sheet.Rows[1:] {

			elements = append(elements, models.Minimumwage{
				Country:     isset(row.Cells, 0)[:len(isset(row.Cells, 0))-1],
				Year:        isset(row.Cells, 1),
				LocalAmount: isset(row.Cells, 2),
				USD:         isset(row.Cells, 3),
			})
		}

	}

	file, err := json.MarshalIndent(elements, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile("./stuff/json/National-Minimum-Wage.json", file, 0644)

	return elements
}

func isset(arr []*xlsx.Cell, index int) string {
	if len(arr) > index {
		return arr[index].String()
	}
	return "NULL"
}
