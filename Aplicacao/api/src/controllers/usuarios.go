package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request){

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil{
		log.Fatal(erro)
	}

	db, erro := banco.Concetar()
	if erro != nil{
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	ususarioID, erro := repositorio.Criar(usuario)
	if erro != nil{
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", ususarioID)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Buscando todos os usuários"))

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Buscando um Usuário!"))

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Atualizando Usuário!"))

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Deletando Usuário!"))

}
