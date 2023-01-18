package enderecos

import "strings"


//TipoEndereco verifica se tem endereço válido e o retorna
func TipoEndereco(endereco string) string{
	tiposValidos := []string{"rua","avenida", "estrada","rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos{
		if tipo == primeiraPalavraDoEndereco{
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tipo Inválido"
}