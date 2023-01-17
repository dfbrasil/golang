package main

import "fmt"

func main(){

	fmt.Println("Maps")

	usuario := map[string]string{

		"nome": "Daniel",
		"sobrenome": "Brasil",
	}
	fmt.Println(usuario)

	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	//MAP aninhado

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro":"João",
			"ultimo":"Pedro",
		},
		"curso":{
			"nome":"Agronomia",
			"campus":"Pici",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2,"nome")
	fmt.Println(usuario2)

	usuario2["centro"] = map[string]string{
		"nome":"CCA",
	}

	fmt.Println(usuario2)
}