package main

import "fmt"

	func main(){
			
			tarefas := make(chan int, 45)

			resultados := make (chan int, 45)

			go worker(tarefas, resultados) //Usado para diminuir filas grandes de tarefa para serem executadas
			go worker(tarefas, resultados) //varios processo fazendo execuções independentes
			go worker(tarefas, resultados)
			go worker(tarefas, resultados)
		

			for i := 0; i <45; i++{
				tarefas <- i
			}

			close(tarefas)

			for i:= 0; i<45; i++{
				resultado := <-resultados
				fmt.Println((resultado))

			}
		}

	func fibonacci(posicao int) int {
		if posicao <=1{
			return posicao
		}
		return fibonacci(posicao - 2) + fibonacci(posicao - 1)
	}
	
	
	func worker(tarefas <-chan int, resultados chan<- int){ // recebe e envia dados respectivamente
		for numero := range tarefas{
			resultados <- fibonacci(numero)
		}
	}