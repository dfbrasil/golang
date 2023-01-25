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
