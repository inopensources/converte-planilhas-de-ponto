package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/********************************
*informações banco de dados local
*********************************/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "conversor_teste"
)

/********************************
* Realizando conexão com o banco
*********************************/
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
}


func formatString(word string) string{

	word = strings.ToUpper(word)

	//removendo acentuação
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	word, _, _ = transform.String(t, word)

	return word
}

/****************************
* Retorna id de funcionário
******************************/
func getIdEmployee(name string) int{

	/*------------------------------making connection-----------------------------------
	 *todo: seria ideal utilizar a funçao connect | referencia a db e err se perdendo
	 */

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
	/*-----------------------------------------------------------------------------------*/

	name = formatString(name)

	var id_user int
	sqlStatement := `SELECT id FROM usuario WHERE nome LIKE $1`
	err = db.QueryRow(sqlStatement, name).Scan(&id_user)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}

	return id_user
}

/****************************
* Retorna id de natureza
******************************/
func getIdNatureza(descricao_natureza string) int{

	/*------------------------------making connection-----------------------------------
		*todo: seria ideal utilizar a funçao connect | referencia a db e err se perdendo
		*/

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
	/*-----------------------------------------------------------------------------------*/


	descricao_natureza = strings.ToLower(descricao_natureza)
	//capitalizing the first letter of each word
	descricao_natureza = strings.Title(descricao_natureza)

	var id_natureza int
	sqlStatment := `SELECT id FROM tipos_ausencia WHERE natureza LIKE $1`
	err = db.QueryRow(sqlStatment, descricao_natureza).Scan(&id_natureza)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}

	return id_natureza
}

/****************************
* Insere pontos em banco
******************************/
func insertPoints(usuario_id int, gerente_id int, entrada_1 string, saida1 string , entrada2 string, saida2 string , entrada3 string , saida3 string , flg_gerado bool, flg_autorizado_pelo_rh bool, created_at string, observacao string, tipo_ausencia int){

	/*------------------------------making connection-----------------------------------
	*todo: seria ideal utilizar a funçao connect | referencia a db e err se perdendo
	*/

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
	/*-----------------------------------------------------------------------------------*/


	/*****************************************tratando entradas****************************
	* todo: formatar no formato: "2013-12-13 13:13:13" e modularizar
	*/

	entrada_1_formated, _ := strconv.ParseFloat(entrada_1, 64)
	entrada_1_formated *=24

	saida1_formated, _ := strconv.ParseFloat(saida1, 64)
	saida1_formated *=24

	entrada2_formated, _ := strconv.ParseFloat(entrada2, 64)
	entrada2_formated *=24

	saida2_formated, _ := strconv.ParseFloat(saida2, 64)
	saida2_formated *=24

/*	//formatando entradas
	entrada_string := fmt.Sprintf("%f", entrada_1)
	fmt.Println("entrada_string", entrada_string)
	entrada_elements_time := strings.Split(entrada_string, ".")
	fmt.Println("entrada_elements_time", entrada_elements_time)
	entrada_hour := entrada_elements_time[0]
	fmt.Println("entrada_hour", entrada_hour)
	entrada_minutes := entrada_elements_time[1]
	fmt.Println("entrada_minutes", entrada_minutes)
	entrada_1_formated := entrada_hour + ":" + strconv.Itoa(int(entrada_minutes[0]))
	fmt.Println("entrada_1_formated", entrada_1_formated)
*/



	//todo: tratar sql injection aqui
	sqlStatement := `
	INSERT INTO ponto (usuario_id, gerente_id, entrada1, saida1, entrada2, saida2, entrada3, saida3, flg_gerado, flg_autorizado_pelo_rh, created_at, observacao, tipo_ausencia)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
_, err = db.Exec(sqlStatement, usuario_id, gerente_id, entrada_1, saida1, entrada2, saida2, entrada3, saida3, flg_gerado, flg_autorizado_pelo_rh, created_at, observacao, tipo_ausencia)
	if err != nil {
		panic(err)
	}

}




