package main

import "fmt"

func getFullName() (string, string) {
	return "Akhnaf", "Ryan"
}

func main() {
	// firstname, lastname := getFullName()
	// fmt.Println("Hello", firstname, lastname)

	firstname, _ := getFullName()
	fmt.Println("Hello", firstname)
}