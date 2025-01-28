package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado'%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris")
		esperado := "Ola, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Ola mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Ola, mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
