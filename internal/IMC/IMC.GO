package IMC

import "fmt"

type DadosPessoais struct {
	Nome   string
	Idade  int
	Peso   float64
	Altura float64
}

func ColetarEExibirDados() {
	var dados DadosPessoais

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&dados.Nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&dados.Idade)

	fmt.Print("Digite seu peso (kg): ")
	fmt.Scanln(&dados.Peso)

	fmt.Print("Digite sua altura (m): ")
	fmt.Scanln(&dados.Altura)

	// Exibir dados
	fmt.Println("\n--- Dados Informados ---")
	fmt.Println("Nome:", dados.Nome)
	fmt.Println("Idade:", dados.Idade)
	fmt.Println("Peso:", dados.Peso)
	fmt.Println("Altura:", dados.Altura)

	// Calcular IMC
	imc := dados.Peso / (dados.Altura * dados.Altura)
	fmt.Printf("IMC: %.2f\n", imc)

	// Verificar maioridade
	maiorDeIdade := dados.Idade >= 18
	fmt.Println("É maior de idade?", maiorDeIdade)
}
