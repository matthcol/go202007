package main

import (
	"calcul"
	"fmt"
)

func main() {
	var a uint = 15
	var b uint = 12
	p := calcul.Pgcd(a, b)
	fmt.Printf("Le pgcd de %v et %v est %v\n", a, b, p)
}
