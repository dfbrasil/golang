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

// inserir t.Parallel() no início de cada função fará que os testes rodem em paralelo
// go test --cover verifica o quanto o teste cobre o código
// go test --coverprofile cobertura.txt salva o resultado do teste em um txt
// go tool cover --func=resultado.txt vai iterpretar no terminal o resultado do teste.txt criado
// go tool cover --html=resultado.txt vai iterpretar no terminal o resultado do teste.txt criado e dizer quais linhas estão com problema ou não coberto pelo teste