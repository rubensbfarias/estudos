package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tip float inferido pelo compilador

	//forma reduzida de criar um VAR

	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)
}
