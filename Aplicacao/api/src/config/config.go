package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	// StringConexaoBanco é a string de conexão com o Mysql
	StrinConexaoBanco = ""

	//Porta onda a API estará rodando
	Porta = 0

)

//Carregar as variáveis de ambiente dentro da aplicação
func Carregar(){ //não vai receber parametros nem retornar nada, pois mexe com variaveis que estão fora dela e estão disponiveis par a API inteira

	//vai usar o pacote godotenv para ler os arquivos que estão no .env

	var erro error

	if erro = godotenv.Load(); erro != nil{
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil{
		Porta = 9000
	}

	//Montar a String de conexao com o banco

	StrinConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	os.Getenv("DB_USUARIO"),
	os.Getenv("DB_SENHA"),
	os.Getenv("DB_NOME"),
)
}