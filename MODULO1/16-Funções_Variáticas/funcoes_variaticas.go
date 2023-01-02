package main

import "fmt"


func soma(numeros ... int) int{
	total := 0
	for _, numero := range(numeros){
		total += numero
	}
	return total
}

func escrever(texto string, numeros ... int){ //Só pode um variático por função e obrigatoriamente tem que ser o último
	for _,numero := range numeros{
		fmt.Println(texto,numero)
	}
}


func main() {

	fmt.Println("Funções variáticas recebem N parâmetros")
	fmt.Println(soma(1,2,3,4,5,6,6))

	escrever("olá mundo", 2,6,6,8)
}