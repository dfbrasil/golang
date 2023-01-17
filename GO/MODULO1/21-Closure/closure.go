package main

import "fmt"

func closure() func(){
	texto := "Dentro da função closure"

	funcao := func ()  {
		fmt.Println(texto)	
	}

	return funcao
}

func main(){
	fmt.Println("Função closure") //Referenciam variáveis que estão fora do seu corpo
	texto := "dentro da função main"
	fmt.Println(texto)

	funcao_nova := closure() //usa a variável quando foi declarada
	funcao_nova()

}