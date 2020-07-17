package cities

import "fmt"

// City represents a city with name, zipcode and its population
type City struct {
	Name       string
	Zipcode    string
	Population uint
}

// UpdatePopulation : function updating population of a city
func UpdatePopulation(pc *City, newPopulation uint) {
	pc.Population = newPopulation
}

// SetPopulation : method updating population of a city
func (c *City) SetPopulation(newPopulation uint) {
	c.Population = newPopulation
}

// String method convert City to string
func (c City) String() string {
	return fmt.Sprintf("%v %v (%v p)", c.Zipcode, c.Name, c.Population)
}

// String2 method convert City to string (method 2)
func (c City) String2() string {
	return fmt.Sprintf("<n:%v z:%v p:%v>", c.Name, c.Zipcode, c.Population)
}

// FilterCities filters cities with zipcode
// Several cities can share the same zipcode
func FilterCities(citiesIn []City, zipcode string) []City {
	res := make([]City, 0, 3)
	for _, c := range citiesIn {
		if c.Zipcode == zipcode {
			res = append(res, c)
		}
	}
	return res
}
