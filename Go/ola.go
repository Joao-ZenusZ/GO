package main

import "fmt"

const prefixoOlaPortugues = "Ola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	if idioma == "Espanhol" {
		return "Hola, " + nome
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("", ""))
}
