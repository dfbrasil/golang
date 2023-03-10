package servidor

import (
	"crud/banco1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct{
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}
	
//CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		w.Write([]byte("Falha ao ler o corpo da requisção"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		w.Write([]byte("Erro ao converter usuário para struct"))
	}
	
	//Abrindo a conexão com o banco
	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	//INSERINDO usuário no banco
	//Prepare statment
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil{
		w.Write([]byte("Erro ao criar o statment"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil{
		w.Write([]byte("Erro ao executar o statment"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil{
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	//Status codes
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
//fim da função CriarUsuario
}

//BuscarUsuarios traz todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){

//abrir conexão com banco de dados
	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

//comando que vai trazer todos os usuarios
//SELECT * FROM USUARIOS
//não se usa o coamndo prepare, usa-se o comando Query, que retorna linhas da tabela
	linhas, erro := db.Query("select * from usuarios") //case insensitive //Fez uma query que retorna linhas
	if erro != nil{
		w.Write([]byte("Erro ao buscas os usuarios!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario//slice de usuarios
	for linhas.Next(){//aqui se intera pelas linhas do banco para o método ir montando a busca
		var usuario usuario //só um usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil{//scanea cada um das informações da linha em ordem e joga nas prop do usuário
			w.Write([]byte("Erro ao escanar o usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)

	//Transformar o slice de usuários em JSON
	//não será feito com Marshal nem Unmarshal
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil{
		w.Write([]byte("Erro ao converter usuarios para JSON"))
	}
//fim da função BuscarUsuários
}

//BuscarUsuario traz um usário específico salvo no banco dedados
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
//primeiro se lê o parâmentro que está sendo passado na rota
//usa uma função do pacote Mux que retorna o parâmentro da requisição

	parametros := mux.Vars(r) // vom como string e tem que se converter para número inteiro

	ID, erro := strconv.ParseUint(parametros["id"],10, 32)// a conversão com essa função são 3 parametros, 1: a variável que se quer converter, 2: a base decimal, e 3: o tamanho do dado, no caso 32 bits

	if erro != nil{
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	//abrir concexão com o banco

	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("Erro ao concetera oa banco de dados"))
		return
	}

	//fazer select e passar a ID para dentro dele

	linha, erro := db.Query("select * from usuarios where id = ?", ID)

	if erro != nil{
		w.Write([]byte("Erro buscar o usuario"))
		return
	}

	//criar usuário , não como um slice de usuario

	var usuario usuario

	if linha.Next(){
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil{
			w.Write([]byte("Erro scanear o usuario"))
		return
		}
	}

	// fazer o esquema do encoder

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil{
		w.Write([]byte("Erro ao converter o usuario para JSON"))
		return
	}
//fim da função buscar usuario
}

//AtualizarUsuario atualiza os dados de um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"],10,32)
	if erro != nil{
		w.Write([]byte("Erro ao ler parametro ID"))
		return
	}
//LER corpo da requisição e abrir conexão com o banco

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		w.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("Erro ao conectar banco de dados"))
		return
	}

	//FAZ O statment para operações que não seja de consulta

	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil{
		w.Write([]byte("Erro ao crirar o statemt"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil{
		w.Write([]byte("Erro ao atualizar usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
//fim da função atualizar usuario
}

//DeletarUsuario deleta os dados de um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
//Ler os parâmentros
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10,32)
	if erro != nil{
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	//abrir concexão

	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("Erro ao concetar o banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil{
		w.Write([]byte("Erro ao criar o statment"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil{
		w.Write([]byte("Erro ao deletar o usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}