package main

import "fmt"

type Ville struct {
	nom        string
	cp         uint32
	nbHabitant uint32
}

func main() {
	villes := []Ville{
		{"Toulouse", 31000, 450000},
		{"Fonsorbes", 31470, 11000},
	}
	fmt.Println("Before:", villes)
	updatePopulation(&villes[0], 450023)
	updatePopulation(&villes[1], 10999)
	fmt.Println("After:", villes)

	ville := Ville{"नई दिल्ली", 110, 250000}
	fmt.Println("Before:", ville)
	updatePopulation(&ville, 250001)
	fmt.Println("After:", ville)
}

func updatePopulation(pv *Ville, newPopulation uint32) {
	pv.nbHabitant = newPopulation
	fmt.Println("Inside function:", pv)
}

func incrementAllPopulation(villes []Ville, increment uint32) {

}
