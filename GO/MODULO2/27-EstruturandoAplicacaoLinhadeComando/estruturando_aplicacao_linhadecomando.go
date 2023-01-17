package main

import (
	"fmt"
	"linha-de-comando/app"
	"os"
	"log"
)

func main(){

	fmt.Println("Estruturando a aplicação de Linha de Comando")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}