package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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

	if erro = usuario.Preparar("cadastro"); erro != nil{ //passa cadastro aqui pois o método prepara precisa de um parâmetro. E quando a requisição chegar no método de validar ele vai ver que a etapa é de cadastro e vai verificar se a senha está em branco.
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
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"],10,64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//Aqui será visto se o usuário tem permissão para alterar o valor, se o usuario 1 pode mudar o valor do usuario 1 aula 92
	usuarioIDNotoken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNotoken{
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuario que não seja o seu "))
		return
	}

	
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil{
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
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request){

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"], 10,64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNotoken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNotoken{
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar um usuario que não seja o seu "))
		return
	}

	db, erro := banco.Concetar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//SeguirUsuario permite que um usario siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request){
//primeira coisa a ser feita é extrarir o usuarioId do token. Quem vai seguir é o usuário que está fazendo a requisição e quem vai ser seguido é o usuário que está no parâmetro.

	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"],10,64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID{
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))//o errors.New() sobrecreve a mensagem de erro padrão do http.StatusXXX
		return
	}

	db, erro := banco.Concetar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//PararDeSeguirUsuario permite que um usuario deixe de seguir outro
func PararDeSeguir(w http.ResponseWriter, r *http.Request){
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	//Ler o usuadrioID que está na reuisição
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID{
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível parar de seguir vc mesmo"))
		return
	}

	db, erro := banco.Concetar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()


	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//BuscarSeguidores trás todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request){

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

	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

//BuscarSeguindo traz todos os usuarios que um determinado usuario está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request){

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

	usuarios, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

//AtualizarSenha atualiza a senha do usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request){
	usuarioIDNotoken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioid"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioIDNotoken != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar a senha de um usuário que não seja o seu"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	var senha modelos.Senha
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil{
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

	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com a senha que está salva no banco"))
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}