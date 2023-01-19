package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct{ //struct com o campos abaixo mapeados para o JSON
	Nome string `json:"nome"` // representa a chave dentro do JSON quando for convertido
	Raca string  `json:"raca"`
	Idade uint `json:"idade"` //a chave para o sjon não é obrigatório ser a mesma que está no struct
}

func main(){

	cachorroEmJSON := `{"nome":"Rex","raca":"PitBull","idade":3}`

	var c cachorro
	//o json.Unmarshal precisa receber um slice de bytes

	//precisa tratar pois o Unmarshal retorna só um valor que é um ERRO, tem que tratar o erro
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil{//precisa passar a referencia de memória & pois se quer que a variável se altere.
		log.Fatal(erro)
	}


	fmt.Println(c) //formato de Struct

	// transformando em MAP

	cachorro2EmJSON :=  `{"nome":"Tal","raca":"Pé duro"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil{//precisa passar a referencia de memória & pois se quer que a variável se altere.
		log.Fatal(erro)
	}

	fmt.Println(c2)

}