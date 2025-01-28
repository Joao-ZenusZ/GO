package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	verificacaoResultadoEsperado := func(t *testing.T, repeticoes, esperado string) {
		t.Helper()
		if repeticoes != esperado {
			t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
		}
	}
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"
	verificacaoResultadoEsperado(t, repeticoes, esperado)
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	repeticao := Repetir("c", 10)
	fmt.Println(repeticao)
	// Output: cccccccccc
}
