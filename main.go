package main

import (
	"database/sql"
	"fmt"
	"github.com/extrame/xls"
	_ "github.com/lib/pq"
	"log"
	"strconv"
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
	entrada_1 float64
	entrada_2 float64
	entrada_3 float64
	entrada_4 float64
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
				roww_nome_funcionario := sheet.Row(4)
				var nome_funcionario = roww_nome_funcionario.Col(2)

				//creating structure employee todo: pegar id através de nome
				emp := Employee{
					name: nome_funcionario,
					id:       95,
				}

				fmt.Println("------------------------------------")

				fmt.Println(nome_funcionario)
				fmt.Println("Employee", emp)

				fmt.Println("------------------------------------")

				row_index := 1
				for row_index <= int(num_row) {

					row := sheet.Row(row_index)



					if row.Col(0) == "F O L H A D E F R E Q U E N C I A" {

						//movendo ponteiro para iniciar  leitura de horários
						row_index += 3

						for row.Col(6) != " "{

							points_initial_row := sheet.Row(row_index)

							entrada_1, _ := strconv.ParseFloat(points_initial_row.Col(1), 64)
							entrada_2, _ := strconv.ParseFloat(points_initial_row.Col(2), 64)
							entrada_3, _ := strconv.ParseFloat(points_initial_row.Col(3), 64)
							entrada_4, _ := strconv.ParseFloat(points_initial_row.Col(4), 64)

							natureza := points_initial_row.Col(6)
                            point := Point{
								data: "12-12-12",
								entrada_1: entrada_1*24,
							    entrada_2: entrada_2*24,
							    entrada_3: entrada_3*24,
							    entrada_4: entrada_4*24,
								natureza: natureza,
								natureza_id: 1,
							}

							fmt.Print(entrada_1*24, "  ")
							fmt.Print(entrada_2*24, "  ")
							fmt.Print(entrada_3*24, "  ")
							fmt.Print(entrada_4*24, "  ")
							fmt.Println(natureza)

							fmt.Println("Point", point)

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

					//percorre colunas dentro de linha
					for col_index := row.FirstCol(); col_index  < row.LastCol(); col_index ++ {
						fmt.Print("[",row_index,":",col_index,"]>", row.Col(col_index ), "  ")
					}

					fmt.Println()
				}
			}

		}
	}
}

/********************************************************************************************************
* from queries.go todo: move to there
 **********************************************************************************************************/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "conversor_teste"
)


func connect(){

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")



	/****************************
	* FROM METHOD getUserIdFromName
	*todo: move to there !!!
	******************************/

	var id_user int

	err = db.QueryRow(`SELECT id FROM usuario WHERE nome LIKE 'HENRIQUE'`).Scan(&id_user)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id_user)


	/****************************
	* FROM METHOD getIdFromNatureza
	*todo: move to there !!!
	******************************/

	var id_natureza int

	err = db.QueryRow(`SELECT id FROM tipos_ausencia WHERE natureza LIKE 'Normal'`).Scan(&id_natureza)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id_natureza)


	/****************************
	* FROM METHOD inserting points
	*todo: move to there  !!!
	******************************/
    //todo: tratar sql injection aqui
	/*sqlStatement = `
	INSERT INTO ponto (usuario_id, gerente_id, entrada1, saida1, entrada2, saida2, entrada3, saida3, flg_gerado, flg_autorizado_pelo_rh, created_at, observacao, tipo_ausencia)
	VALUES (95, 40,'2013-12-13 13:13:13','2013-12-13 13:13:13','2013-12-13 13:13:13','2013-12-13 13:13:13','2013-12-13 13:13:13','2013-12-13 13:13:13', false, false, '2013-12-13 13:13:13', 'nada', 4263)`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)

	}*/


}

func main(){


	print("=======================\n" +
		"Conversor Licença XLS/XSLX\n" +
		"========================\n")

	connect()

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
	//getAllContent()

}
