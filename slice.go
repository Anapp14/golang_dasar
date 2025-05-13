package main

import "fmt"

func main() {
	names := [...]string{"Anap", "Budi", "Caca", "Dede", "Eka", "Fani", "Gina", "Hana", "Ika", "Joko"}

	slice1 := names[4:8]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	var slice4 []string = names[2:5]
	fmt.Println(slice4)
}