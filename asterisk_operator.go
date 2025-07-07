package main 

import "fmt"

type Address struct {
	City, Provience, Country string
}

func main() {
	address1 := Address{"Yogyakarta" , "DIY" , "IDN"} 
	address2 := &address1 
	address2.City = "Bantul"
	fmt.Println(address1)
	fmt.Println(address2)

	//address2 = &Address{"Sleman" , "Jabar", "IDN"}
	*address2 = Address{"Sleman" , "Jabar", "IDN"}
	fmt.Println(address1)
	fmt.Println(address2)

}