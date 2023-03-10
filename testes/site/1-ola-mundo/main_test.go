package main

import "testing"

func TestOla(t *testing.T) {
	
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Daniel")
		esperado := "Olá Daniel"

		if resultado != esperado{
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Diz Olá Mundo quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá mundo"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})
}