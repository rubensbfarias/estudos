package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		numero1 int
		numero2 int
	)

	// Solicita ao usuário que digite um número inteiro
	fmt.Println("Digite um número inteiro:")
	fmt.Scanln(&numero1)

	// Verifica se o número é positivo, negativo ou zero
	if numero1 > 0 {
		// Calcula a raiz quadrada do número e exibe o resultado na tela
		raiz := math.Sqrt(float64(numero1))
		fmt.Printf("O número é positivo e sua raiz quadrada é: %.2f\n", raiz)
	} else if numero1 < 0 {
		// Exibe uma mensagem indicando que não é possível calcular a raiz quadrada de um número negativo
		fmt.Println("O número é negativo e não é possível calcular sua raiz quadrada.")
	} else {
		// Exibe uma mensagem indicando que o número é zero
		fmt.Println("O número é zero.")
	}

	// Verifica se o número é um número primo
	if numero1 > 1 {
		primo := true
		for i := 2; i <= int(math.Sqrt(float64(numero1))); i++ {
			if numero1%i == 0 {
				primo = false
				break
			}
		}
		if primo {
			fmt.Println("O número é primo.")
		} else {
			fmt.Println("O número não é primo.")
		}
	} else {
		fmt.Println("O número não é primo.")
	}

	// Solicita ao usuário que digite um segundo número inteiro
	fmt.Println("Digite outro número inteiro:")
	fmt.Scanln(&numero2)

	// Verifica se o primeiro número é divisível pelo segundo número
	if numero2 == 0 {
		fmt.Println("Não é possível dividir por zero.")
	} else if numero1%numero2 == 0 {
		fmt.Printf("%d é divisível por %d.\n", numero1, numero2)
	} else {
		fmt.Printf("%d não é divisível por %d.\n", numero1, numero2)
	}
}
