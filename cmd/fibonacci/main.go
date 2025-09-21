package main

import (
	"fmt"

	"github.com/Adrianotorres411/meu-projeto-go/internal/fibonacci"
)

func main() {
	fmt.Println("Testando pacote Fibonacci:")

	// Teste básico
	fmt.Printf("F(10) = %d\n", fibonacci.Fibonacci(10))

	// Sequência completa
	seq := fibonacci.Sequence(10)
	fmt.Printf("Sequência: %v\n", seq)

	// Usando a função de impressão
	fibonacci.PrintSequence(7)
}
