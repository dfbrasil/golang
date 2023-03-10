package main

import "fmt"

const NumeroInteracoes int = 5

func Repetir(char string) string{
	var i int
	var repeticoes string
	fmt.Println("Digite a quantidade de Interações")
	fmt.Scan(NumeroInteracoes)
	for i = 0; i<NumeroInteracoes; i++{
		repeticoes += "d"
	}
	return repeticoes
}

