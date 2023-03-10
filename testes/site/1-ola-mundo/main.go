package main

import "fmt"

const prefixoOlaPortugues = "Olá "

func Ola(nome string) string{
	return prefixoOlaPortugues + nome
}

func main(){
	fmt.Println(Ola("teste"))
}
