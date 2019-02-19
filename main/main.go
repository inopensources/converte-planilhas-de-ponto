package main

import (

	"../../github.com/360EntSecGroup-Skylar/excelize"
	"../src/github/tealeg/xlsx"
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

	/*//getting a single cell
	getCell("Sheet1", "B1", "../samples/file_example_XLSX_10.xlsx")

	//getting all rows
	getAllRows("Sheet1", "../samples/file_example_XLSX_10.xlsx")*/


	//getting a single cell
	//getCell("Sheet1", "B1", "../samples/file_example_XLSX_10.xlsx")
	//getting all rows
	//getAllRows("Sheet1", "../samples/file_example_XLS_10.xls")

	//using extrame library
/*	if xlFile, err := Open("Table.xls", "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			fmt.Print("Total Lines ", sheet1.MaxRow, sheet1.Name)
			col1 := sheet1.Rows[0].Cols[0]
			col2 := sheet1.Rows[0].Cols[0]
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Rows[uint16(i)]
				col1 = row1.Cols[0]
				col2 = row1.Cols[11]
				fmt.Print("\n", col1.String(xlFile), ",", col2.String(xlFile))
			}
		}
	}*/

	//using tealeg library
	excelFileName := "../samples/file_example_XLS_10.xls"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Print(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}


}