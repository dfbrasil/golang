package main

import (
	"api/src/config"
	"api/src/router"
	// "crypto/rand"
	// "encoding/base64"
	"fmt"
	"log"
	"net/http"
)
//função que gera inicialmente a chave randomica
// func init(){
// 	chave := make([]byte, 64)//chave secrete randomica do token

// 	if _, erro := rand.Read(chave); erro != nil{
// 		log.Fatal()
// 	}

// 	//converter essa chave mestring para inserir no .env

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)


// }

func main(){

	config.Carregar()

	r := router.Gerar()

	// fmt.Println(config.SecretKey)//Testando a secretkey aula 89

	fmt.Printf("Escutando na porta %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}