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

// type MyWebServerType bool

// type login int
// type welcome int

// func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w, "Página de Login")
// }

// func (wc welcome) ServeHTTP (w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w, "Página de Boas-vindas")
// }

func myLogging(w http.ResponseWriter, r *http.Request){ //ServeHTTP implenta um tipo de Handler
	// fmt.Fprintln(w, "Hello mothefuckers") //w -> Responsewriter
	// fmt.Fprintf(w, "Request is: %+v", r)
	// fmt.Fprintf(w, `
	// <!DOCTYPE html>
	// 	<html lang="en">
	// 	<head>
	// 		<meta charset="UTF-8">
	// 		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	// 		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	// 		<title>Teste de Login</title>
	// 	</head>
	// 	<body>
	// 		<h1>
	// 			Insira seu nome e senha
	// 		</h1>
	// 	</body>
	// 	</html>
	// `)
	if r.Method == "GET"{
		fmt.Fprintln(w, "using GET for logging endpoint")
	}
	if r.Method == "POST"{
		fmt.Fprintln(w, "Usando o método post na request do endpoint")
	}
	fmt.Fprintln(w, "Página de Login")
}

func myWelcome(w http.ResponseWriter, r *http.Request){ //ServeHTTP implenta um tipo de Handler
	// fmt.Fprintln(w, "Hello mothefuckers") //w -> Responsewriter
	// fmt.Fprintf(w, "Request is: %+v", r)
	// fmt.Fprintf(w, `
	// <!DOCTYPE html>
	// 	<html lang="en">
	// 	<head>
	// 		<meta charset="UTF-8">
	// 		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	// 		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	// 		<title>Teste de Boas Vindas</title>
	// 	</head>
	// 	<body>
	// 		<h1>
	// 			Bem - Vindos
	// 		</h1>
	// 	</body>
	// 	</html>
	// `)

	fmt.Fprintln(w, "Página de Boas-vindas")
}


func main(){
	// var k MyWebServerType //O handler implementado acima está sendo passado como segundo parâmentro da função ListenAndServe
	http.HandleFunc("/login", myLogging)
	http.HandleFunc("/welcome", myWelcome)

	// http.Handle("/login", http.HandlerFunc(myLogging)) //este aceita um handler e não uma função, como se envia um handle aqui? Usando HandlerFunc, que parece ser um método mas não é. HandlerFunc é um Type. O que tem que ser fazer é passar a ação (ou função). por um type, ele também pode ter métodos, no caso Response e Request
	// http.Handle("/welcome", http.HandlerFunc(myWelcome))

	// var i login
	// var j welcome
	// http.Handle("/login", i)
	// http.Handle("/welcome", j)

	fmt.Println("Escutando na porta 8080")
	http.ListenAndServe("localhost:8080", nil)
}


//Três maneiras de chamar uma handle

//1. HandleFunca recebe o path e a função que tem a mesma assinatura, mas não tem o mesmo nome
//2. Handle aceita recebe a url e o handler
//3. o handler poderia ser um handler puro ou um um handler convertido usando um handler func type 