package main

import "fmt"

// Função nomeada
func saudacaoNomeada(nome string) string {
	return "Olá " + nome
}

func main() {
	// Usando a função nomeada
	fmt.Println(saudacaoNomeada("Adriano"))

	// Criando uma função anônima e atribuindo a uma variável
	saudacaoAnonima := func(nome string) string {
		return "Olá " + nome
	}

	// Usando a função anônima
	fmt.Println(saudacaoAnonima("lorranny"))
}
