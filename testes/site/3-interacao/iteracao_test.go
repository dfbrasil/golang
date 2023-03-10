package main

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("d")
	esperado := "ddddd"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas retornou '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repetir("d")
    }
}