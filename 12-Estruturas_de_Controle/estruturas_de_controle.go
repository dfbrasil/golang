package main

import "fmt"

func main()  {
	
	fmt.Println("Estruras de Controle")

	numero := -10

	if numero>15{
		fmt.Println("O número é maior que 15")
	} else{
		fmt.Println("O número é menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0{ //If Init
		fmt.Println("Númeor é maior que 0")
	} else if numero < -10 {
		fmt.Println("Número é menor que 0")
	} else {
		fmt.Println("O número está entre -10 e 10")
	}
}