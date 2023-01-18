package main

import "fmt"

func main(){

	fmt.Println("Canais com buffer")

	canal := make(chan string, 2) //capacidade de buffer 2 ,se colocar um terceiro valor para o canal, dará deadlock

	canal <- "Olá mundo"
	canal <- "Programando em GO"

	mensagem := <-canal
	mensagem2 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}