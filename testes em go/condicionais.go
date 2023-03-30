// OPERADORES LÓGICOS
// E = &&
// OU = ||
// IGUAL = ==
// DIFERENTE = !=

//CRIAR UM SISTEMA DE CNH ONDE O PROGRAMA IRA CAPTURAR A IDADE E O NOME DA PESSOA
//E IRÁ REALIZAR A SEGUINTE LÓGICA
//SE A PESSOA TIVER +18 ANOS ELA PODE TER A CARTEIRA DE MOTORISTA
//SE ELA TIVER -18 ELA NÃO PODERÁ TER A CARTEIRA DE MOROTIRISTA
//SE ELA TIVER +65 ANOS ELA NÃO PODE DIRIGIR
//CASO A PESSOA NÃO TENHA A IDADE PARA DIRIGIR IRÁ APARECER
//QUANTO TEMPO FALTA PARA QUE ELA CONSIGA TIRAR A CARTEIRA.

package main

import "fmt"

func main() {

	var nome string
	var sobreNome string
	var idade int
	var resultado int
	var idadeMinima int

	fmt.Println("Qual o seu nome?")
	fmt.Scanln(&nome)
	fmt.Println("Qual o seu sobrenome?")
	fmt.Scanln(&sobreNome)
	fmt.Println("Qual a sua idade?")
	fmt.Scanln(&idade)

	if idade > 65 {
		fmt.Println("Não pode dirigir")
	} else if idade < 18 {
		idadeMinima = 18
		resultado = (idadeMinima - idade)
		fmt.Println("Falta:", resultado)
		fmt.Println("Você não pode tirar a carteira de motorista!")

	} else if idade >= 18 {
		fmt.Println("Pode tirar a carteira de motorista")
	}
}
