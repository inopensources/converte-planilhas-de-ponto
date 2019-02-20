package main

import (
	"fmt"

	"../src/xls"
)


/****************************
* Methods for .xlsx extension
*****************************/

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


/****************************
* Methods for .xls extension
****************************/

func getQuantEmployers(pathFile string) int {

	var num_employers = 0

	if xlFile, err := xls.Open(pathFile, "utf-8"); err == nil {

		if err != nil {
			fmt.Print(err)
		}else{
			num_employers = xlFile.NumSheets()-1;
			fmt.Print("Número de funcionários:\n", num_employers , "\n") //desconsidera aba modelo
	     }
	}

	return num_employers
}

func getNameEmployer() string{
 return ""
}


func mountDate(date string){
	fmt.Println(date)
}

func getDate(pathFile string, indexSheet uint16) string {
	return ""
}

func getPoints(){


	if xlFile, err := xls.Open("../samples/CartaoPonto_Sistemas_2017.xls", "utf-8"); err == nil {

		if err != nil {
			fmt.Print(err)
		}

		for i := 0; i <= xlFile.NumSheets(); i++ {

			//verifica se é aba modelo
			if xlFile.GetSheet(i).Name != "MODELO" {

				sheet := xlFile.GetSheet(i)
				num_row := sheet.MaxRow

				//retorna nome de usuário:
				roww_nome_funcionario := sheet.Row(4)
				var nome_funcionario = roww_nome_funcionario.Col(2)
				fmt.Println(nome_funcionario)

				row_index := 1
				for row_index <= int(num_row) {

					row := sheet.Row(row_index)

					if row.Col(0) == "F O L H A D E F R E Q U E N C I A" {

						//movendo ponteiro para iniciar  leitura de horários
						row_index += 3

						for row.Col(6) != " " {

							points_initial_row := sheet.Row(row_index)
							fmt.Print(points_initial_row.Col(1), "  ")
							fmt.Print(points_initial_row.Col(2), "  ")
							fmt.Print(points_initial_row.Col(3), "  ")
							fmt.Println(points_initial_row.Col(4))
							row_index += 1

						}

					}

					row_index++
				}

			}

		}


	}
}

func getAllContent() {

	//using extrame library
	if xlFile, err := xls.Open("../samples/CartaoPonto_Sistemas_2017.xls", "utf-8"); err == nil {

		if err != nil {
			fmt.Print(err)
		}

		fmt.Print("Número de funcionários:\n", xlFile.NumSheets()-1, "\n") //desconsidera aba modelo

		for i := 0; i < xlFile.NumSheets(); i++ {

			//verifica se é aba modelo
			if xlFile.GetSheet(i).Name != "MODELO" {

				fmt.Println("Funcionário: ", xlFile.GetSheet(i).Name)
				sheet := xlFile.GetSheet(i)

				//rettorna quantidades de linhas em aba
				num_row := sheet.MaxRow
				//fmt.Print("Total Lines:\n", sheet.MaxRow, "\n")

				/*Percorre todos os elementos da aba*/
				//percorre linhas dentro de aba
				for row_index := 1; row_index <= int(num_row); row_index++ {

					row := sheet.Row(row_index)

					//percorre colunas dentro de linha
					for index := row.FirstCol(); index < row.LastCol(); index++ {
						//fmt.Println("coluna ", index, ">", row.Col(index))

					}
				}
			}

		}
	}
}

func main(){

	print("=======================\n" +
		"Conversor Licença XLS/XSLX\n" +
		"========================\n")


	/*Using excelize library!*/
	/*//getting a single cell
	getCell("Sheet1", "B1", "../samples/file_example_XLSX_10.xlsx")

	//getting all rows
	getAllRows("Sheet1", "../samples/file_example_XLSX_10.xlsx")*/

	//getting a single cell
	//getCell("Sheet1", "B1", "../samples/file_example_XLSX_10.xlsx")
	//getting all rows
	//getAllRows("Sheet1", "../samples/file_example_XLS_10.xls")




}
