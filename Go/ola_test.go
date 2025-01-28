package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Chris")
	esperado := "Ola, Chris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
