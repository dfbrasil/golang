package modelos

import "time"

type Usuario struct{
	ID uint64 `json:"id,omitempty"` //se o campo Id tiver em branco ele não passa para o JSON, ele tira o campo
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}