package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20))

	numbers	:= []int{20, 39, 50, 12}
	fmt.Println(sumAll(numbers...))
	
}