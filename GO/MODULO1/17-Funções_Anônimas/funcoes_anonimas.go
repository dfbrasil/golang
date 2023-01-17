package main

import "fmt"




func main(){
	fmt.Println("Funções Anônimas")

	retorno := func(texte string) string{
		return fmt.Sprintf("Recebido -> %s %d", texte, 22 )
	}("Passando os valores dos parâmetros texte e ")

	fmt.Println(retorno)

}