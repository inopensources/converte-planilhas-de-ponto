package main

import (
	"fmt"
	"github.com/extrame/xls"
	_ "github.com/lib/pq"
)



/**********************************
* Structs todo: move to their place
***********************************/
type Employee struct {
	name string
	id int
}

type Point struct {
	data string
	entrada_1 string
	entrada_2 string
	entrada_3 string
	entrada_4 string
	natureza string
	natureza_id int
}


/****************************
* Methods for .xlsx extension
*****************************/


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


func getPoints(){

	fmt.Println("Getting Points..")


	if xlFile, err := xls.Open("./samples/CartaoPonto_Sistemas_2017.xls", "utf-8"); err == nil {

		if err != nil {
			fmt.Print(err)
		}

		for i := 0; i <= xlFile.NumSheets(); i++ {

			//verifica se é aba modelo
			if xlFile.GetSheet(i).Name != "MODELO" {


				sheet := xlFile.GetSheet(i)
				num_row := sheet.MaxRow

				//retorna nome de usuário:
				row_nome_funcionario := sheet.Row(4)
				var nome_funcionario = row_nome_funcionario.Col(2)
				id_employee := getIdEmployee(nome_funcionario)

				emp := Employee{
					name: nome_funcionario,
					id:       id_employee,
				}

				fmt.Println("------------------------------------")
				fmt.Println("Funcionário", emp)
				fmt.Println("------------------------------------")

				row_index := 1
				for row_index <= int(num_row) {

					row := sheet.Row(row_index)


					if row.Col(0) == "F O L H A D E F R E Q U E N C I A" {

						//movendo ponteiro para iniciar  leitura de horários
						row_index += 3
						row := sheet.Row(row_index)
						condicao_parada := row.Col(6)

						for condicao_parada != " "{

							points_initial_row := sheet.Row(row_index)

							natureza := points_initial_row.Col(6)
							condicao_parada = natureza
							id_natureza := getIdNatureza(natureza)

							point := Point{
								data: "12-12-12",
								entrada_1: points_initial_row.Col(1),
							    entrada_2: points_initial_row.Col(2),
							    entrada_3: points_initial_row.Col(3),
							    entrada_4: points_initial_row.Col(4),
								natureza: natureza,
								natureza_id: id_natureza,
							}


							fmt.Println(points_initial_row.Col(6))
							fmt.Println("Point", point)

                            /*Armazenando Ponto*/
							insertPoints(id_employee, 40, "2013-12-13 13:13:13", "2013-12-13 13:13:13","2013-12-13 13:13:13","2013-12-13 13:13:13","2013-12-13 13:13:13","2013-12-13 13:13:13", false, false, "2013-12-13 13:13:13", "nada", id_natureza)

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
	if xlFile, err := xls.Open("./samples/CartaoPonto_Sistemas_2017.xls", "utf-8"); err == nil {

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

						if row != nil {

							col_index_first := row.FirstCol()
							col_index_last  := row.LastCol()

							//percorre colunas dentro de linha
							for col_index := col_index_first; col_index  < col_index_last; col_index ++ {
								fmt.Print("[",row_index,":",col_index,"]>", row.Col(col_index ), "  ")
							}

							fmt.Println()
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

	//getPoints()
	getAllContent()

	// Output:
	// resume

}
