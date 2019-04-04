package main

import (
	"fmt"
	"github.com/extrame/xls"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
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
	saida_1 string
	entrada_2 string
	saida_2 string
	natureza string
	natureza_id int
	observacao string
	flg_gerado bool
	flg_autorizado_pelo_rh bool
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
	dia := 0
	mes := 0
	ano := ""
	data := ""


	if xlFile, err := xls.Open("./samples/file_example_XLS_10.xls", "utf-8"); err == nil {

		if err != nil {
			fmt.Print(err)
		}

		for i := 0; i <= xlFile.NumSheets(); i++ {

			sheet := xlFile.GetSheet(i)
			mes  = 0

			//verifica se é aba modelo
			if xlFile.GetSheet(i).Name != "MODELO" {

				num_row := sheet.MaxRow

				//retorna nome de usuário:
				row_nome_funcionario := sheet.Row(4)
				var nome_funcionario = row_nome_funcionario.Col(2)
				id_employee := getIdEmployee(nome_funcionario)

				employee := Employee{
					name: nome_funcionario,
					id:       id_employee,
				}

				for row_index := 1; row_index <= int(num_row); row_index++ {

					row := sheet.Row(row_index)

					if row != nil {

						if row.Col(0) == "F O L H A D E F R E Q U E N C I A" {
							/*assumindo uma folha de frequência por mês*/
							mes++
							dia = 0

							//movendo ponteiro para iniciar  leitura de horários
							row_index += 3

							condicao_parada := false //row.Col(6)

							for !condicao_parada {

								/*assumindo cada linha um dia*/
								dia++
								points_initial_row := sheet.Row(row_index)
								natureza := points_initial_row.Col(6)

								if natureza != "" && natureza != " "{

									id_natureza := 12 //getIdNatureza(natureza)
									observacao := points_initial_row.Col(7)

									//usando TrimPrefix para capturar ano em data
									ano = strings.TrimPrefix(ano, "01/01/")
									data = strconv.Itoa(mes) + "-" + strconv.Itoa(dia) +  "-" + ano

									point := Point{
										data:                   data,
										entrada_1:              points_initial_row.Col(1),
										saida_1:                points_initial_row.Col(2),
										entrada_2:              points_initial_row.Col(3),
										saida_2:                points_initial_row.Col(4),
										natureza:               natureza,
										natureza_id:            id_natureza,
										observacao:             observacao,
										flg_gerado:             false,
										flg_autorizado_pelo_rh: false,
									}

									/*Armazenando Ponto no banco de dados*/
									insertPoints(point, employee)

									row_index += 1

								}else {
									condicao_parada = true
								}

								}


							}

						}
					}

				}else{

				//utilizando modelo para recuperar ano
				row_date := sheet.Row(1)
				ano = strings.Split(row_date.Col(8), ";")[0]// everything before the query
			}

		}


	}
}

func getAllContent() {

	//using extrame library
	if xlFile, err := xls.Open("./samples/file_example_XLS_10.xls", "UTF-32"); err == nil {

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
								fmt.Print("[",row_index,":",col_index,"]>", row.Col(col_index), "  ")
							}

							fmt.Println()
						}

				}

			}

		}
	}
}

/*monta estruturas com as datas e pontos de funcinários do mês*/
func mount_objects() {
		if xlFile, err := xls.Open("./samples/CartaoPonto_Sistemas_2018.xls", "utf-8"); err == nil {

			    //pegando aba modelo para recuperar datas
				if sheet1 := xlFile.GetSheet(0); sheet1 != nil {

					for i := 0; i <= (int(sheet1.MaxRow)); i++ {
						row := sheet1.Row(i)

						if row != nil {

							col_index_first := row.FirstCol()
							col_index_last := row.LastCol()

							//percorre colunas dentro de linha
							for col_index := col_index_first; col_index < col_index_last; col_index ++ {
								//fmt.Print("[", i, ":", col_index, "]>", row.Col(col_index), "  ")
								if row.Col(col_index) == "Mês/Ano" {
									i += 1 //andando uma linha abaixo
									row := sheet1.Row(i)
									fmt.Println("[",i,":", col_index,"]",row.Col(8))
								}
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

	getPoints()


}
