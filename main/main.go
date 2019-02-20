package main

import (
	"fmt"

	"../src/xls"
)

/*
*get all the rows in the Sheet indicated
 */
/*func getAllRows(sheet string, pathToFile string) {

	xlsx, err := excelize.OpenFile(pathToFile)
	if err != nil {
		fmt.Print(err)
	}

	rows := xlsx.GetRows(sheet)
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}*/

/**
*get the value from cell by given worksheet name and axis
 */
/*func getCell(sheet string, axis string, pathToFile string) string {

	xlsx, err := excelize.OpenFile(pathToFile)
	if err != nil {
		fmt.Print(err)
	}

	cell := xlsx.GetCellValue(sheet, axis)
	fmt.Println(cell)
	return cell
}*/

func main() {

	print("=======================\n" +
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
	if xlFile, err := xls.Open("../samples/CartaoPonto_Sistemas_2017.xls", "utf-8"); err == nil {

		if err != nil {
			fmt.Print(err)
		}

		fmt.Print("Num Sheets:\n", xlFile.NumSheets(), "\n")

		for i := 0; i < xlFile.NumSheets(); i++ {
			fmt.Println(xlFile.GetSheet(i).Name, "\n")
			sheet := xlFile.GetSheet(i)
			fmt.Print("Total Lines:\n", sheet.MaxRow, "\n")

			for row_index := 1; row_index < 15; row_index++ {
				row := sheet.Row(row_index)
				for index := row.FirstCol(); index < row.LastCol(); index++ {
					fmt.Println(index, "==>", row.Col(index), " ")

				}
			}

		}

		/*if sheet1 := xlFile.GetSheet(0); sheet1 != nil {

			fmt.Print("Sheet Name:\n", sheet1.Name, "\n")

			fmt.Print("Total Lines:\n", sheet1.MaxRow, "\n")

			row := sheet1.Row(0)

			for index := row.FirstCol(); index < row.LastCol(); index++ {
				fmt.Println(index, "==>", row.Col(index), " ")

			}

		}*/
	}

	//using tealeg library
	/*excelFileName := "../samples/file_example_XLS_10.xls"
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
	}*/

}
