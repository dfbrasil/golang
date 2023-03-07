package main

import (
	"fmt"
	"net/http"
)

//Esta interface é o parâmetro Handler (2° Parâmentro do ListenAndServe). E é implementada na documentação da maneira abaixo.
//Esta interface só possui um método em sua implementação que é ServeHTTP
// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

//Implementando um tipo Próprio

type MyWebServerType bool

func (MyWebServerType )ServeHTTP(w http.ResponseWriter, r *http.Request){ //ServeHTTP implenta um tipo de Handler
	fmt.Fprintln(w, "Hello mothefuckers") //w -> Responsewriter
	fmt.Fprintf(w, "Request is: %+v", r)
}


func main(){
	var k MyWebServerType //O handler implementado acima está sendo passado como segundo parâmentro da função ListenAndServe 
	http.ListenAndServe("localhost:8080", k)
}