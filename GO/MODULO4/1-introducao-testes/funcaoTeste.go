package main

import (
	"fmt"
	"modulo1/enderecos"
	
)

func main(){

	tipoEndereco := enderecos.TipoEndereco("Avenida Brasil")
	fmt.Println(tipoEndereco)

}

//rodar o go test -v ele testa todas as funcoes
