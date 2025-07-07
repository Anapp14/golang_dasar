package main

import "fmt"

func main() {
	name := "Akhnaf"

	switch name {
	case "Akhnaf":
		fmt.Println("Halo Akhnaf")
	case "Anap":
		fmt.Println("Kamu Bukan Akhnaf, tapi Anap")
	default:
		fmt.Println("Kamu Siapa?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama kamu terlalu panjang")
	case false:
		fmt.Println("Nama kamu terlalu pendek")
	}
}