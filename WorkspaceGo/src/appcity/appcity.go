package main

import (
	"cities"
	"fmt"
)

func main() {
	c1 := cities.City{Name: "Toulouse", Zipcode: "31000", Population: 45000}
	c2 := cities.City{Name: "Pau", Zipcode: "64000"}
	c3 := cities.City{"Fonsorbes", "31470", 11000}
	fmt.Println("Init:", c1, c2, c3)
	fmt.Println("Explicit call String method:", c2.String())
	fmt.Println("Explicit call String method 2:", c2.String2())
	cities.UpdatePopulation(&c2, 77000)
	fmt.Println("After update by function:", c2)
	c2.SetPopulation(77051)
	fmt.Println("After update by method:", c2)
}
