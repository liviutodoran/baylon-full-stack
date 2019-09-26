package utilstools

import (
	
	"log"
	"github.com/tealeg/xlsx"
	"babylon-stack/api/models"	
)

func GetDataXLX()[]models.MinWage {
	excelFileName := "./stuff/test.xlsx"
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
			/*fmt.Println("\n", row.Cells[0])
			fmt.Println("\n", row.Cells[1])
			fmt.Println("\n", row.Cells[2])
			fmt.Println("\n", row.Cells[3])
			/*for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}*/
		}
	}
	return elements

	
}
