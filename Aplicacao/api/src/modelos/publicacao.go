package modelos

import "time"

//Publicação reperesenta uma publicação feita por um usuario
type Publicacao struct{
	ID uint64 `json:"id,omitempty"`//O omitempty como o nome mostra não exibe o campo caso ele esteja vazio
	Titulo string `json:"titulo:omitempty"`
	Conteudo string `json:"conteudo:omitempty"`
	AutorID uint64 `json:"autorId:omitempty"`
	AutorNick uint64 `json:"autorNick:omitempty"`
	Curtidas uint64 `json:"curtidas"`
	CriadaEm time.Time `json:"criadaEm:omitempty"`
}
