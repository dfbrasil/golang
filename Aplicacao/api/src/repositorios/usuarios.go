package repositorios

import (
	"api/src/modelos"
	"database/sql"
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
