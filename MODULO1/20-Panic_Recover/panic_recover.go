package main

import "fmt"

func recuperar_execucao(){
	if recup := recover(); recup!= nil{ //só executada se o Panic for executado, já que recup será nulo caso não haja Panic
		fmt.Println("Execução recuperada com sucesso")

	}
	
}

func aluno_esta_aprovado(n1 float32, n2 float32) bool{
	defer recuperar_execucao()
	media := (n1 + n2) / 2

	if media >6 {
		return true
	} else if media < 6{
		return false
	}

	panic("A média é exatamente 6") //Mata a execução do programa
	// antes de matar tudo e parar a execução ele chama as DEFER

}

func main(){
	fmt.Println("Panic e Recover")

	fmt.Println(aluno_esta_aprovado(6, 6))

	fmt.Println("Pós execução")



}