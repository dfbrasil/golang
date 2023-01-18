package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal // canal vai esperar receber o valor, recebendo um valor para canal
		if !aberto{
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")

	for mensagem := range canal{
		fmt.Print(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i <5; i++{
		canal <- texto //canal vai estar recebendo o texto, mandando um valor pra dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}