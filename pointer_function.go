package main

import "fmt"

type Address struct {
	City, Provience, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "IDN"
}

func main() {
	var address Address = Address{}
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)
}