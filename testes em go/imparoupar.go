// QUANDO FOR PAR (IDADE)
// QUANDO FOR IMPAR (IDADE)

// AMANHÃ ÀS XX:XX (MATTHAIOS)
package main

import "fmt"

func main() {
	var (
		numeral  int
		restante int
	)

	fmt.Println("Digite o valor de entrada")
	fmt.Scanln(&numeral)

	restante = numeral % 2

	fmt.Println(restante)

	if restante == 0 {
		fmt.Printf("O numero %d é par!", numeral)
	} else {
		fmt.Printf("O numero %d é impar!", numeral)
	}

}
