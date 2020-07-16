package main

import "fmt"

func main() {
	mots := []string{"go", "python", "java", "go", "ruby", "php", "r", "go", "c#", "python"}
	fmt.Printf("%d %v\n", len(mots), mots)
	var res map[string]uint
	// TODO
	fmt.Printf("%d %v\n", len(res), res)
}
