package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) { //Handlefunc precisa receber dois paramentros específicos, o http.Responsewriter e o htt.Request vão fazer o preocesso de response e request
		w.Write([]byte("Olá Mundo!!")) // funcção do ResponseWrite que escreve uma mensagem e tem que estar no formato de slice de byte
	}

func usuarios(w http.ResponseWriter, r *http.Request) { 
		w.Write([]byte("Carregar página de usuário!!")) 
	}	


func main(){
	//Recebe o URI da rota e a função que recebe a requisção se vai saber lidar com essa requisição
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000",nil)) //depois dessa linha progrma trava...


}

// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DE COMUNICAÇÃO

// CLIENTE (FAZ A REQUISIÇÃO )- SERVIDOR (PROCESSA A REQUISIÇÃO E ENVIA A RESPOSTA)

// Request - Response

// Rotas (identificar o tipo de mensagem/processamento que o servidor vai ter que fazer)
	//URI - Identificador do Recurso
	//Método - o que quer fazer com o recurso - GET, POST, PUT, DELETE