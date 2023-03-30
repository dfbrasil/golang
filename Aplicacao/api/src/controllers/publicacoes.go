package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CriarPublicação adiciona uma nova publicação no DB
func CriarPublicacao(w http.ResponseWriter, r *http.Request){
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao modelos.Publicacao
	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	publicacao.AutorID = usuarioID

	if err = publicacao.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Concetar()
	if err != nil{
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

//BuscasPublicacoes traz as publicações que apareceriam no feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request){
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := banco.Concetar()
	if err != nil{
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio:= repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repositorio.Buscar(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

//BuscarPublicacao traz uma única publicacao
func BuscarPublicacao(w http.ResponseWriter, r *http.Request){
	parametro := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}	

	db, err := banco.Concetar()
	if err != nil{
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio:= repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)

}

//AtualizarPublicacao altera os dados de uma publicacao
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request){
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Concetar()
	if err != nil{
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("você não pode alterar uma publicação que não seja sua"))
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao modelos.Publicacao
	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = publicacao.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = repositorio.Atualizar(publicacaoID, publicacao); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//DeletarPublicacao apaga uma publicacao
func DeletarPublicacao(w http.ResponseWriter, r *http.Request){
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Concetar()
	if err != nil{
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("você não pode deletar uma publicação que não seja sua"))
		return
	}

	if err = repositorio.Deletar(publicacaoID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioid"], 10, 64)
	if err != nil{
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Concetar()
	if err != nil{
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repositorio.BuscarPorUsuario(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, publicacoes)
}