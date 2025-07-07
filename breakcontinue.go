package main

import "fmt"

func main() {
	for i := 1; i <10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke-", i)
	}

	for a := 1; a < 10; a++ {
		if a%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke-", a)
	}
}