package main

import "fmt"

func main() {
	var coffee string
	var nbHabitant int
	nbHabitant = 471941
	ville := "Toulouse" // declare and assign (same type as right expression)
	// fmt.Println("Hello, 世界,", ville, "(", nbHabitant, ")")
	fmt.Printf("Hello, 世界, %v (%d hab.), je bois du café <%v>\n", ville, nbHabitant, coffee)
	nbHabitant = 11743
	ville = "Fonsorbes"
	coffee = "long à grain"
	fmt.Printf("Hello, 世界, %v (%d hab.), je bois du café <%v>\n", ville, nbHabitant, coffee)
}
