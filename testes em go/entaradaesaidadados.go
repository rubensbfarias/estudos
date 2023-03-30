// Criar um sistema em que o programa pergunte seu nome e seu sobrenom
// A pessoa vai inserir os dados
// Depois o programa vai retornar os dados dee nome e sobrenome retornar na mesma linha

package main

import "fmt"

func main() {

	var nome string
	var sobreNome string

	fmt.Println("Qual o seu nome?")
	fmt.Scanln(&nome)
	fmt.Println("Qual o seu sobrenome?")
	fmt.Scanln(&sobreNome)

	nomeCompleto := "Seu nome é " + nome + " seu sobrenome é " + sobreNome

	fmt.Println(nomeCompleto)
}
