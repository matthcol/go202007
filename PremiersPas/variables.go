package main

import "fmt"

func main() {
	var ville string = "नई दिल्ली"
	var nbHabitant uint32 = 21700000 // unsigned int 32 bits : 0 to 2^32-1
	var solde int32 = -15000         // signed int 32 bits : -2^31 à 2^31-1
	nbHabitant = 4000000000          // ok en unsigned
	// solde = 4000000000 // nok en signed
	var defaultInt int = 8000000000 // either int32 either int64
	fmt.Printf("%s %d %d %d\n", ville, nbHabitant, solde, defaultInt)
	fmt.Printf("%T\n", defaultInt)

	var temperature float32 = 16.8 // float IEEE754 simple precision
	temperature = 0.1
	fmt.Printf("%0.15f %0.15f %0.15f\n", temperature, 2*temperature, 3*temperature)
	var prix float64 = 0.1
	fmt.Printf("%0.18f %0.18f %0.18f\n", prix, 2*prix, 3*prix)
	var faitIlBeau bool = true
	fmt.Printf("%v %t\n", faitIlBeau, faitIlBeau)
}
