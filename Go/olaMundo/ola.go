package main

import "fmt"

const prefixoOlaPortugues = "Ola, "
const espanhol = "Espanhol"
const prefixoOlaEspanhol = "Hola, "
const frances = "Francês"
const prefixoOlaFrances = "Bonjour, "
const ingles = "Inglês"
const prefixoOlaIngles = "Hello, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("", ""))
}
