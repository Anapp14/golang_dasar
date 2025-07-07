package main

import "fmt"

func getFullName() (firstname, middlename, lastname string) {
	firstname = "Akhnaf"
	middlename = "Ryan"
	lastname = "Ghifari"
	return firstname, middlename, lastname
}

func main() {
	a, b, c := getFullName()
	fmt.Println("Hello", a, b, c)
}