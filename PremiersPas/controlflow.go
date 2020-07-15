package main

import "fmt"

func main() {
	for temperature := 5.6; temperature < 30; temperature += 10 {
		fmt.Printf("* Aujourd'hui, il fait %.1f\n", temperature)
		if temperature < 10 {
			fmt.Println("Veste, bonnet et gants")
		} else if temperature < 20 { // 10 <= temperature <20
			fmt.Println("Un petit gilet")
		} else {
			fmt.Println("Tenue d'été")
		}
		fmt.Println("Maintenant je peux sortir")
	}
	temperature := 5.6
	for temperature < 30 {
		fmt.Printf("+ Aujourd'hui, il fait %.1f\n", temperature)
		temperature += 10
	}
	for ; temperature < 40; temperature++ {
		fmt.Printf("~ Aujourd'hui, il fait %.1f\n", temperature)
	}

	switch choix := 1; choix {
	case 1:
		fmt.Println("Choix 1")
	case 2:
		fmt.Println("Choix 2")
	case 3:
		fmt.Println("Choix 3")
	default:
		fmt.Println("Choix défaut")
	}

}
