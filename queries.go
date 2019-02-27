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


func remove_accent(word string) string{

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	word, _, _ = transform.String(t, word)

	return word
}

func mount_timestamp(time_entrada string, date string) string{

	timestamp_formated := ""

	time_entrada_float, _:= strconv.ParseFloat(time_entrada, 64)
	time_entrada_float *=24.0

	time_entrada_string := fmt.Sprintf("%f", time_entrada_float)
	time_entrada_elements := strings.Split(time_entrada_string, ".")

	entrada_hour := time_entrada_elements[0]
	entrada_minutes := time_entrada_elements[1]
	entrada_minutes = entrada_minutes[:2]

	entrada_minutes_int, _:= strconv.Atoi(entrada_minutes)
	if entrada_minutes_int > 60 {
		entrada_minutes_int = 59
	}
	entrada_minutes_string := fmt.Sprintf("%d", entrada_minutes_int)


	timestamp_formated = date + " " + entrada_hour + ":" + entrada_minutes_string + ":" + "00"

	return timestamp_formated
}


func create_user(name string) int{

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

	//todo: tratar sql injection aqui
	sqlStatement := `
	INSERT INTO usuario(nome, senha, email, acesso, ativo, excluido, sistema, foto, digital, carga_horaria)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, name,' ', ' ', ' ', true, false, false, ' ', ' ', ' ').Scan(&id)

	if err != nil {
		panic(err)
	}

    return id
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

	name = strings.ToUpper(name)
	name = remove_accent(name)
	fmt.Println(name)

	var id_user int

	sqlStatement := `SELECT id FROM usuario WHERE nome LIKE $1`
	err = db.QueryRow(sqlStatement, name).Scan(&id_user)
	if err == sql.ErrNoRows {
		//log.Fatal("Usuário não encontrado | Criando usuário")
		fmt.Println("Usuário não encontrado | Criando usuário")
		//se usuário não existir, cria usuário
		id_user = create_user(name)

	}
	if err != nil {
		//log.Fatal(err)
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

    /*formatando descrição da natureza*/
	descricao_natureza = strings.ToLower(descricao_natureza)
	//capitalizing the first letter of each word
	descricao_natureza = strings.Title(descricao_natureza)
	descricao_natureza = remove_accent(descricao_natureza)

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
func insertPoints(point Point, employee Employee){

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


	/*------------------------------------tratando entradas--------------------------------
	* todo: formatar no formato: "2013-12-13 13:13:13" e modularizar
	*/

	//fmt.Println(point, employee)

    //entradas:
	usuario_id := employee.id
	gerente_id := 40

	date := point.data

	//entradas em format string e com apenas as datas
	entrada_1 := mount_timestamp(point.entrada_1, date)
	saida_1 := mount_timestamp(point.saida_1, date)// point.saida_1
	entrada_2 := mount_timestamp(point.entrada_2, date) //point.entrada_2
	saida_2 := mount_timestamp(point.saida_2, date) //point.saida_2
	entrada_3 := mount_timestamp(point.entrada_2, date)
	saida_3 :=  mount_timestamp(point.saida_2, date)

	flg_gerado := point.flg_gerado
	flg_autorizado_pelo_rh := point.flg_autorizado_pelo_rh

	created_at := entrada_1
	observacao := point.observacao
	tipo_ausencia := point.natureza_id

  /*----------------------------------------------------------------------------------------------------------*/

	sqlStatement := `
	INSERT INTO ponto (usuario_id, gerente_id, entrada1, saida1, entrada2, saida2, entrada3, saida3, flg_gerado, flg_autorizado_pelo_rh, created_at, observacao, tipo_ausencia)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
_, err = db.Exec(sqlStatement, usuario_id, gerente_id, entrada_1, saida_1, entrada_2, saida_2, entrada_3, saida_3, flg_gerado, flg_autorizado_pelo_rh, created_at, observacao, tipo_ausencia)
	if err != nil {
		panic(err)
	}

}




