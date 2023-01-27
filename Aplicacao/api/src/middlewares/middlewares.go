package middlewares

import (
	"fmt"
	"log"
	"net/http"
)


//Logger Escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//Autentitcar verifica se o usuario fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc{//func (w http.ResponseWriter, r *http.Request)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")//vai chamar a função de validação do token e vai executar o que veio no parametro, o next
		next(w, r)
	}

}