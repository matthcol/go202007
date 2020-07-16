package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // time de maintenant de type Type
	fmt.Println(t)
	fmt.Printf("%v %v %v %v %v\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	y, m, d := t.Date()
	fmt.Printf("%v %v %v\n", y, m, d)
	fmt.Println(t.Format("01/02/2006 à 3h4'5"))
}
