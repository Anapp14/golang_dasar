package main

import "fmt"

func main13 () {
	var names[3]string

	names[0] = "akhnaf"
	names[1] = "ryan"
	names[2] = "ghifari"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array langsung
	var values = [...]int {
		90,
		100,
		110,
		120,
		130,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[3] = 140
	fmt.Println(values)
}