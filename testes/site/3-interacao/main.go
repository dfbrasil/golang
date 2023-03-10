package main

import "fmt"


func main()  {
	var NumeroInteracoes int
	var i int
	var repeticoes string
	fmt.Println("Digite a quantidade de Interações")
	fmt.Scanln(&NumeroInteracoes)
	for i = 0; i<NumeroInteracoes; i++{
		repeticoes += "d"
	}
	fmt.Println(repeticoes)
	
}