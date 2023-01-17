package main

import "fmt"


func calculos_matematicos(n1,n2 int) (soma int, sub int){
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main(){

	fmt.Println("Funções com retorno nomeado")

	fmt.Println(calculos_matematicos(2,3))

	soma, sub := calculos_matematicos(10,20)
	fmt.Println(soma, " e ", sub)
}