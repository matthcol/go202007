package main

import "fmt"

type Ville struct {
	nom        string
	cp         uint32
	nbHabitant uint32
}

func main() {
	var ville Ville
	fmt.Println(ville)
	fmt.Printf("%d <%v>, %d habitants\n", ville.cp, ville.nom, ville.nbHabitant)

	ville2 := Ville{"Toulouse", 31000, 450000}
	fmt.Println(ville2)
	fmt.Printf("%d <%v>, %d habitants\n", ville2.cp, ville2.nom, ville2.nbHabitant)

	ville2 = Ville{"Fonsorbes", 31470, 11000}

	villes := []Ville{
		{"Toulouse", 31000, 450000},
		{"Fonsorbes", 31470, 11000},
	}
	fmt.Println(villes)
	villes = append(villes, Ville{"Pau", 64000, 77000})
	fmt.Println(villes)

	for i, v := range villes {
		fmt.Printf("- %d - %v (%d)\n", i, v, v.cp)
	}

	// not good for update
	ville = villes[0] // copy
	fmt.Printf("Ville n°0: %v\n", ville)
	ville.nbHabitant++
	fmt.Printf("Villes: %v ; ville n°0: %v\n", villes, ville)

	var pville *Ville
	fmt.Printf("default pointer value to nil ? %t\n", pville == nil)
	pville = &villes[0]
	//fmt.Printf("Ville n°0 par pointer: %v %d\n", *pville, (*pville).nbHabitant)
	fmt.Printf("Ville n°0 par pointer: %v %d\n", *pville, pville.nbHabitant)
	pville.nbHabitant++
	//(*pville).nbHabitant++
	fmt.Printf("Après modif via pointer %x %v %v\n", pville, *pville, villes[0])

	// maj avec loop : not good at all !!!!
	for _, v := range villes {
		v.nbHabitant++ // ++ on copy
	}
	fmt.Println("Villes après ++ nb habitants: ", villes)

	// maj avec loop : the right solution
	for i := range villes {
		villes[i].nbHabitant++ // ++ in the slice
	}
	fmt.Println("Villes après ++ nb habitants on all slice: ", villes)
}
