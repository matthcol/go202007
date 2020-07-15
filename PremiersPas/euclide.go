// algorithme d'Euclide
package main

import "fmt"

func main() {
	a := 15
	b := 12
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	fmt.Printf("Le pgcd est %v\n", a)
}
