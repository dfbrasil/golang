package main

import "fmt"

func funcao1(){
	fmt.Println("Executando a função 1")
}

func funcao2(){
	fmt.Println("Executando a função 2")
}

func aluno_esta_aprovado(nota1 float32, nota2 float32) bool{
	defer fmt.Println("Média calculada. O resultado será retornado") //será executado antes do return da função
	fmt.Println("Entrando na função que verifica se o aluno está aprovado")

	media := (nota1 + nota2) / 2

	if media >= 6{
		return true
	} else{
		return false
	}
		

}

func main(){
	fmt.Println("Defer") //Adia a execução de um comando (pedaço de código) pra antes da última execução
	// muito executado em BD

	funcao1()
	funcao2()

	defer funcao1() //DEFER = ADIAR
	funcao2()

	fmt.Println(aluno_esta_aprovado(5,5))

}