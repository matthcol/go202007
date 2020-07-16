// algorithme d'Euclide
package main

import "fmt"

func main() {
	var a uint = 15
	var b uint = 12
	p := pgcd(a, b)
	fmt.Printf("Le pgcd de %d et %d est %d\n", a, b, p)
}

func pgcd(a, b uint) uint {
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	return a
}
