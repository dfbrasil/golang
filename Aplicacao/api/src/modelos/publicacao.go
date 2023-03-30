package modelos

import (
	"errors"
	"strings"
	"time"
)

//Publicação reperesenta uma publicação feita por um usuario
type Publicacao struct{
	ID uint64 `json:"id,omitempty"`//O omitempty como o nome mostra não exibe o campo caso ele esteja vazio
	Titulo string `json:"titulo,omitempty"`
	Conteudo string `json:"conteudo,omitempty"`
	AutorID uint64 `json:"autorId,omitempty"`
	AutorNick string `json:"autorNick,omitempty"`
	Curtidas uint64 `json:"curtidas"`
	CriadaEm time.Time `json:"criadaEm,omitempty"`
}

//Preparar vai chamara funções que vai formatar e validar os dados da publicação
func (publicacao *Publicacao) Preparar() error{
	if erro := publicacao.validar(); erro != nil{
		return erro
	}
	publicacao.formatar()
	return nil
}

// validar vai verificar se o titulo e o conteudo da publicação estão preenchidos
func (publicacao *Publicacao) validar() error{
	if publicacao.Titulo == ""{
		return errors.New("o título é obrigatório e não pode estar em branco")
	}
	if publicacao.Conteudo == ""{
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}
	return nil
}

//formatar vai chamar funções que vai formatar os dados da publicação
func (publicacao *Publicacao) formatar(){
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
