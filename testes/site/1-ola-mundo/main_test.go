package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Mundo!")
	esperado := "Olá, Mundo!"

	if resultado != esperado {
		t.Errorf("esperado '%s', resultado '%s'", resultado, esperado)
	}
}