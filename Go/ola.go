package main

import "fmt"

func Ola(nome string) string {
	return "Ola, " + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
