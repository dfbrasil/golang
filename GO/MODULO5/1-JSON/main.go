package main

import (
	"bytes"
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

	c := cachorro{"Rex","PitBull", 3}
	
	fmt.Println(c)

	//json.Marshal() MAP para JSON ou STRUCT para JSON o json.Unmarshal faz o inverso

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil{
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON) //ate aqui retornaria um slite de bytes [123 34 110 111 109 101 34 58 34 82 101 120 34 44 34 114 97 99 97 34 58 34 80 105 116 66 117 108 108 34 125]

	saida := bytes.NewBuffer(cachorroEmJSON) // Transforma esses bytes um JSON propriamente dito

	fmt.Println(saida)

}