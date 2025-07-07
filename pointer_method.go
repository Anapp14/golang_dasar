package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	akhnaf := Man{"Akhnaf"}
	akhnaf.Married()

	fmt.Println(akhnaf.Name)
}