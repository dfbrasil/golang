package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Login responsável por autenticar o usuario na API
func Login(w http.ResponseWriter, r *http.Request){
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario//usuario está vindo na requisição
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
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
	usuarioSalvoNobanco, erro := repositorio.BuscarPorEmail(usuario.Email)//usuario que está salvo no banco
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNobanco.Senha, usuario.Senha); erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNobanco.ID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))// esse_ é pra ignorar o erro, por enquanto
}