package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T){

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia BR116", "Rodovia"},
		{"Estrada do Fio", "Estrada"},
		{"Praça", "Tipo Inválido"},
	}


	for _, cenario := range cenariosDeTeste{
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado{
				t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", retornoRecebido,
				cenario.retornoEsperado)
		}
	}
}