package main

import "fmt"

func main9 () {
	type NoKTP string

	var ktpAnap NoKTP = "1111111"

	var contoh string =  "22222222"
	var contohKtp NoKTP = NoKTP(contoh)
	
	fmt.Println(ktpAnap)
	fmt.Println(contohKtp)
}