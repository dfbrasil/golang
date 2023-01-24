package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da API
type Rota struct{

	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool

}

//Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router{ //vai receber um router dewscofigurado e vai retornar a router configurada, usando o HandleFunc
	rotas := rotasUsuarios

	for _, rota := range rotas{
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)//URI , Método e Funçao que será executada
	}

	return r

}