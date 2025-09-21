package main

import (
	"fmt"
)

func main() {
	var nome string
	var idade int
	var peso, altura float64

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	fmt.Print("Digite seu peso (kg): ")
	fmt.Scanln(&peso)

	fmt.Print("Digite sua altura (m): ")
	fmt.Scanln(&altura)

	fmt.Println("\n--- Dados Informados ---")
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Peso:", peso)
	fmt.Println("Altura:", altura)

	imc := peso / (altura * altura)
	fmt.Printf("IMC: %.2f\n", imc)

	maiorDeIdade := idade >= 18
	fmt.Println("É maior de idade?", maiorDeIdade)
}
