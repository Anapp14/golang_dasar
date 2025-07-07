package main

import "fmt"

func main() {
	name := "Ana"

	if name == "Akhnaf" {
		fmt.Println("Halo Akhnaf")
	} else if name == "Anap"{
		fmt.Println("Kamu Bukan Akhnaf, tapi Anap")
	} else {
		fmt.Println("Kamu Siapa?")
	}


	if length := len(name); length > 5 {
		fmt.Println("Nama kamu terlalu panjang")
	} else if length < 5 {
		fmt.Println("Nama kamu terlalu pendek")
	} else {
		fmt.Println("Nama kamu pas")
	}

}