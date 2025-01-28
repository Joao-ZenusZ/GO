package main

import "fmt"

const prefixoOlaPortugues = "Ola, "

func Ola(nome string) string {
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
