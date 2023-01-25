package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"strconv"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request){

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Concetar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){

	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Concetar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuarios ,erro := repositorio.Buscar(nomeOuNick)

	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Concetar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorId(usuarioID)

	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Atualizando Usuário!"))

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Deletando Usuário!"))

}
