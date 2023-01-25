package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)


//JSON essa funçao de resposta genérica, vai receber um statusCode passado, vai colcoar esse status code no header, vai pegr os dados genéricos e transformar para Json. Precisa do ResponseWriter, pois é ele que vai dar a resposta.
//Json retorna uma resposta em Json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}){

	w.Header().Set("Content.Type","application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil{
		log.Fatal(erro)
	}
}


//Erro retorna um erro em formato JSOn
//Erro também vai receber um ResponseWriter, também recebe o StatusCode e o erro propriamente dito
//essa função vai chamar a função JSON basicamente
func Erro(w http.ResponseWriter, statusCode int, erro error){
	JSON(w, statusCode, struct{
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}