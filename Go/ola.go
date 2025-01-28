package main

import "fmt"

const prefixoOlaPortugues = "Ola, "

func Ola(nome string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola(""))
}
