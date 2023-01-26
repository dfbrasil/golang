package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct{
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios{
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error){
	statement, erro := repositorio.db.Prepare("insert into usuarios(nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil{
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil{
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil{
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
	
}


//Buscar traz todos os usuários que atendem um filtro de de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error){
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //está assim //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil{
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next(){
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil{
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil
		
}

//BuscarPorId
func (repositorio usuarios) BuscarPorId(ID uint64) (modelos.Usuario, error){
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil{
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next(){
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil{
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

//Atualizar usuario no danco de dados
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _,erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil{
		return erro
	}// Esse ID é o do parâmetro

	return nil
}

//Deletar usuário do banco de dados

func (repositorio usuarios) Deletar(ID uint64) error{
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _,erro = statement.Exec(ID); erro != nil{
		return erro
	}// Esse ID é o do parâmetro

	return nil

}


//BuscarporEmail busca por email e retorna o seu id e senha com hash
func (repositorio usuarios) BuscarPorEmail (email string) (modelos.Usuario, error){
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil{
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next(){
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil{
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}
