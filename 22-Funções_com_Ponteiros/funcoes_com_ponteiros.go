package main

import (
	"fmt"

)

func inverter_sinal(numero int) int{
	return numero *-1
}


func inverter_sinal_com_ponteiro(numero *int){
	*numero = *numero * -1
}

func main(){
	fmt.Println("Funções com ponteiros")
	
	numero := 20
	fmt.Println(inverter_sinal(numero))

	fmt.Println(numero)

	numero2 := 40

	inverter_sinal_com_ponteiro(&numero2) //TEM que passar o endereço de memória para a função --> &

	fmt.Println(numero2)


}