package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T)  {
	assert.True(t, true, "true is true")
}

func TestSoma(t *testing.T) {
	
	testes := []struct {
		Valores []int
		Resultado int
	}{
		{Valores: []int{1, 2, 3}, Resultado: 6},
        {Valores: []int{1, 2, 3, 4}, Resultado: 10},
	}

	for _, teste := range testes {
        total := Soma(teste.Valores...)

		assert.Equal(t, total, teste.Resultado, "Valor esperado: %d - Valor retornado: %d", teste.Resultado, total)
    }
}

func TestJustPanic(t *testing.T) {
    assert := assert.New(t)
    assert.Panics(func() { JustPanic(true) })
    assert.NotPanics(func() { JustPanic(false) })
}
