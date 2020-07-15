// array, slice, map, struct
package main

import "fmt"

func main() {
	var temperatures [5]float32
	fmt.Println(temperatures)
	var pressions [4]float32
	pressions[0] = 23.4
	pressions[1] = 34.6
	pressions[2] = -12.4
	pressions[3] = 234.12
	fmt.Println(pressions)
	fmt.Println("Pression 0:", 3*pressions[0]+1)
	fmt.Printf("Taille: %d; capacité: %d\n", len(pressions), cap(pressions))
	var somePressions []float32 = pressions[1:3]
	fmt.Printf("Slice de pressions : %v %d %d\n", somePressions, len(somePressions), cap(somePressions))
	somePressions = pressions[3:]
	fmt.Printf("Slice de pressions : %v %d %d\n", somePressions, len(somePressions), cap(somePressions))
	somePressions = pressions[:2]
	fmt.Printf("Slice de pressions : %v %d %d\n", somePressions, len(somePressions), cap(somePressions))
	somePressions = pressions[4:]
	fmt.Printf("Slice de pressions : %v %d %d\n", somePressions, len(somePressions), cap(somePressions))
	// somePressions = pressions[4:6]  // out of bounds for 4-element array
	// fmt.Printf("Slice de pressions : %v %d %d\n", somePressions, len(somePressions), cap(somePressions))

	vitesses := []float32{49.4, 83.3, 145.2}
	// var vitesses []float32 = []float32{49.4, 83.3, 145.2}
	fmt.Printf("Slice de vitesses : %v %d %d\n", vitesses, len(vitesses), cap(vitesses))
	vitesses = append(vitesses, 180.7)
	fmt.Printf("Slice de vitesses : %v %d %d\n", vitesses, len(vitesses), cap(vitesses))
	vitesses = append(vitesses, 280.3)
	fmt.Printf("Slice de vitesses : %v %d %d\n", vitesses, len(vitesses), cap(vitesses))
	vitesses = append(vitesses, 450.0, 500.2, 800.3, 5000.2)
	fmt.Printf("Slice de vitesses : %v %d %d\n", vitesses, len(vitesses), cap(vitesses))

	//anyPressions := append(pressions, 1.1, 2.2, 3.3) // impossible sur un array
	//fmt.Printf("Slice de vitesses : %v %d %d\n", anyPressions, len(anyPressions), cap(anyPressions))

	otherVitesses := make([]float32, 0, 1000) // length 0 capacity 1000
	fmt.Printf("Slice de vitesses : %v %d %d\n", otherVitesses, len(otherVitesses), cap(otherVitesses))
	for v := 90.0; v < 1000.0; v++ {
		otherVitesses = append(otherVitesses, float32(v))
	}
	fmt.Printf("Slice de vitesses : %d %d\n", len(otherVitesses), cap(otherVitesses))
}
