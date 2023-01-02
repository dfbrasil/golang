package main

import "fmt"


var n int


func init(){
	fmt.Println("Executando antes da função init") //Será executada antes da função main
	// pode ser usada par inicializar variávels gloabias por exemplo
	//pode ser usadado para inicair uma configuração ou setup
	n = 10
}

func main() {

	fmt.Println("Função Init")

	fmt.Println("Função main sendo executada !")
	fmt.Println(n)



}