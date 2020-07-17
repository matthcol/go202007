package main

import (
	"cities"
	"fmt"
)

func main() {
	cities.ReadCities("data/cities31.txt")
	fmt.Printf("\n*****************\n")
	cities.ReadCities2("data/cities31.txt")
	fmt.Printf("\n*****************\n")
	cities.ReadCities3("data/cities31.txt")
	fmt.Printf("\n*****************\n")
	citiesLoaded := cities.ReadCities4("data/cities31.txt")
	fmt.Println(citiesLoaded)
	cities31470 := cities.FilterCities(citiesLoaded, "31470")
	fmt.Println(cities31470)
}
