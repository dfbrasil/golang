package main

import "fmt"

type usuario struct{
	nome string
	idade uint8
	endereco endereco
}

type endereco struct{
	logradouro string
	numero uint8
}

func main(){

	fmt.Println("Arquivo Struct")

	var u usuario
	u.nome = "Daniel"
	u.idade = 21
	fmt.Println(u)

	end1 := endereco{"rua tal", 22}

	u2 := usuario{"isis",8, end1}
	fmt.Println(u2)

	u3 := usuario{nome: "Isaac"}
	fmt.Println(u3)

	u4 := usuario{idade: 3}
	fmt.Println(u4)


}