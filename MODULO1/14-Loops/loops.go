package main

import (
	"fmt"
	// "time"
)

func main(){

	fmt.Println("Loops")

	// i:= 0

	// for i < 10{
		
	// 	i++
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Second)
	// }

	// for i := 0; i < 10; i+=2 {
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"dnaiel","freitas","brasil"}

	// for indice, valor := range(nomes){
	// 	fmt.Println(indice,valor)

	// }

	// for _, valor := range(nomes){
	// 	fmt.Println(valor)
	// }

	// for _, letra := range"PALAVRA"{
	// 	fmt.Println(string(letra))
	// }
			
	usuario := map[string]string{
		"nome":"Daniel",
		"Sobrenome":"Brasil",
	}
	for chave, valor := range usuario{
		fmt.Println(chave,valor)
	}

}