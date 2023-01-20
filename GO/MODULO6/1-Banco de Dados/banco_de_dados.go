package main

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/go-sql-driver/mysql" //o _ é pra importar implicitamente o pacote
)

func main(){

	stringConexao := "golang:golang@/estudosgo?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close() //fechando a conexão

	if erro = db.Ping(); erro != nil{ //testa a conexão com o banco
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}