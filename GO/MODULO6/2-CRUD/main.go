package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){

	//CREATE - POST
	//READ - GET
	//UPDATE - PUT
	//DELETE - DELETE


//depois de inserir a função lá em servidor (Adicionar, buscar, etc) é adicionado um router aqui na main

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)//Inicio da aula de insercão de dados
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)//rota para a busca de usuarios, pode ter um mesmo URI mas com métodos diferentes. O método é composto pela URI e pelo método CRUD
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)//rota para busca o usuário específico, no GO se coloca os valores entre {} para dizer que é um valor variável.
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

		

}