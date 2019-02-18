package main


import (
	"../src/github/360EntSecGroup-Skylar/excelize"
	"fmt"
)


/*
*get all the rows in the Sheet indicated
*/
func getAllRows(sheet string, pathToFile string){


	xlsx, err := excelize.OpenFile(pathToFile)
	if err != nil{
		fmt.Print(err)
	}

	rows := xlsx.GetRows(sheet)
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

/**
*get the value from cell by given worksheet name and axis
*/
func getCell(sheet string ,axis string, pathToFile string) string {

	xlsx, err := excelize.OpenFile(pathToFile)
	if err != nil{
		fmt.Print(err)
	}

	cell := xlsx.GetCellValue(sheet, axis)
	fmt.Println(cell)
	return cell
}



func main() {
	print("=======================\n"+
		         "Conversor Licença XSLX\n" +
		         "========================\n")

	//getting a single cell
	getCell("Sheet1", "B1", "../samples/file_example_XLSX_10.xlsx")

	//getting all rows
	getAllRows("Sheet1", "../samples/file_example_XLSX_10.xlsx")

}