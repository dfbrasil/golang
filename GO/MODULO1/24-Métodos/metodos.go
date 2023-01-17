package main

import "fmt"


type usuario struct{
	nome string
	idade uint8
}

func (u usuario) salvar(){ 
	// esse primeiro ( ) é pra indicar à qual função ou estrutura esse método está se referindo
	// esse U é um nome de uma variável qualquer, o que tem que ser igual na chamanda do método é o segundo nome, no caso usuario
	//O método está linkado com alguma estrutura, não é algo solto
	//pode ter retorno, pode passar parâmetro


	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n",u.nome)

}
func (u usuario) maior_de_idade() bool{
	return u.idade>=18
}

func (u *usuario) fazer_aniversario() { // um método com um ponteiro para alterar a idade do usuário.
	//sem o ponteiro o usuario iria alterar, mas apenas dentro do método e não seria alterado em outras funções, como o no método main por exemplo.
	u.idade++
}


func main(){

	fmt.Println("Métodos")

	usuario1 := usuario{"Daniel", 40}

	fmt.Println(usuario1)

	usuario1.salvar()

	usuario2 := usuario{"Isaac",3}

	usuario2.salvar()

	fmt.Println(usuario1.maior_de_idade())
	fmt.Println(usuario2.maior_de_idade())

	usuario1.fazer_aniversario()
	fmt.Println(usuario1.idade)

}