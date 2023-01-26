package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (usuario *Usuario) Preparar(etapa string) error{ //quanod chamar o método vai ver se está numa etapa de cadastro ou etapa de edição de usuario
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil{
		return erro
	}
	
	return nil
}

func(usuario *Usuario) validar(etapa string) error{
	if usuario.Nome == ""{
		return errors.New("o nome é obigatório e não pode estar em branco")
	}

	if usuario.Nick == ""{
		return errors.New("o nick é obigatório e não pode estar em branco")
	}

	if usuario.Email == ""{
		return errors.New("o email é obigatório e não pode estar em branco")
	}

	//pacote que faz a validação do email
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil{
		return errors.New("o e-mail inserido é inválido")
	}
	
	

	if etapa == "cadastro" && usuario.Senha == ""{
		return errors.New("a senha é obigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error{
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil{
			return erro
		}

		usuario.Senha = string(senhaComHash)

	}

	return nil
}