package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	//pegar 2 o mais canais e juntar num so
	canal := multiplexar(escrever("Ola mundo"), escrever("Programando em GO"))

	for i := 0; i<10; i++{
		fmt.Println(<-canal)
	}

}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string{ //canais que só recebem valores, neste caso

	canalDeSaida := make(chan string)

	go func ()  {
		for {
			select{
			case mensagem := <- canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaida <- mensagem
			}

		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string{
	canal := make(chan string)

	go func(){
		for {
			canal <- fmt.Sprintf("VAlor recebido %s", texto) //encapsula uma goroutime e retorna um canal de comunicação
			time.Sleep(time.Millisecond + time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}