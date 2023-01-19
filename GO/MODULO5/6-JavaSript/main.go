package main

import (
	"fmt"
	"html/template" //cria templates baseados no HTML e deixa dinamico
	"log"
	"net/http"
)


type usuario struct{
	Nome string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) { 
		//executa um template específico e recebe 3 parâmetros. O 1° é o reponseWriter, o 2° é o nome do template, o 3° é um dado para passar para esse template seria nesse local a passagem
		u := usuario{
			"daniel",
			"danieldfb@gmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	}

var templates *template.Template //vai conter todos os templates dentro da aplicação




func main(){

	templates = template.Must(template.ParseGlob("*.html"))// está jogando na variavel templates todos os templates que estão sendo criados, referenciando todo mundo que seja .html
	
	http.HandleFunc("/home", home)


	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000",nil)) //depois dessa linha progrma trava...

}
