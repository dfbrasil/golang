package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct{
	ID uint64 `json:"id,omitempty"` //se o campo Id tiver em branco ele não passa para o JSON, ele tira o campo
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}


//Preparar vai chamar os métodos para vaidar e formatar o usuario recebido
func (usuario *Usuario) Preparar() error{
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func(usuario *Usuario) validar() error{
	if usuario.Nome == ""{
		return errors.New("o nome é obigatório e não pode estar em branco")
	}

	if usuario.Nick == ""{
		return errors.New("o nick é obigatório e não pode estar em branco")
	}

	if usuario.Email == ""{
		return errors.New("o email é obigatório e não pode estar em branco")
	}

	if usuario.Senha == ""{
		return errors.New("a senha é obigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(){
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}