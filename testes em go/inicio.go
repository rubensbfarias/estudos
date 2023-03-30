// O que é uma variavel?
// R: Alocação de espaço em memoria
// EX: 4 Vagas,
// Primeira vaga = o carro PRETO, segunda vaga = o carro vermelho
// Terceira verde, quarta amarelo.

// Variaveis

// carros
// vaga1 = preto
// vaga2 = vermelho
// vaga3 = verde
// vaga4 = amarelo

// Tipos de dados

// int = inteiro   (números inteiros)
// double = decimal (valores decimais) obs: até 20 casas decimais
// float = decimal (valores decimais) obs: até 6 casas decimais
// char = caracteres (apenas para caracteres)

package main

import "fmt"

func main() {

	nome := "Rubens"
	sobrenome := "Farias"
	nomeCompleto := "Meu nome é " + nome + " meu sobrenome é " + sobrenome

	// fmt.Println("Meu nome é", nome)
	// fmt.Println("Meu sobrenome é", sobrenome)
	// fmt.Print("Meu nome é ", nome)
	//fmt.Print("Meu sobrenome é ", sobrenome)
	//fmt.Print("Meu nome é ", "Meu sobrenome é ", nome, sobrenome)
	// fmt.Print("Meu nome e meu "+"sobrenome é ", nome+sobrenome)

	fmt.Println(nomeCompleto)
}
