package main

import "fmt"

type Customer struct {
	Name, Address string
	Age			  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var ryn Customer
	fmt.Println(ryn)

	ryn.Name = "Akhnaf Ryan Ghifari"
	ryn.Address = "Yogyakarta"
	ryn.Age = 18
	fmt.Println(ryn)
	fmt.Println(ryn.Name)
	fmt.Println(ryn.Address)
	fmt.Println(ryn.Age)

	desi := Customer {
		Name: "Desi Dewi Anggtaini",
		Address: "Sleman",
		Age: 18,
	}
	fmt.Println(desi)

	anap := Customer{"Anap", "Sleman", 10}
	
	fmt.Println(anap)
	anap.sayHello("ryan")
}