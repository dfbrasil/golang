package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)//2 goroutines que le tem que esperar terminar

	go func(){
		escrever("Olá mundo!!")//goroutine
		waitGroup.Done() //tira 1 do add(2)
	}()

	go func(){
		escrever("Programando em Go")
		waitGroup.Done() // tira -1 do add(2)
	}()
	
	waitGroup.Wait() //se tirar isso ele excuta as duas funções e finaliza o programa
}

func escrever(texto string) {
	for i := 0; i <5; i++{
		fmt.Println(texto) 
		time.Sleep(time.Second)
	}
}