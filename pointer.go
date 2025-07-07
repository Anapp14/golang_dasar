package main
 
import "fmt"

type Address struct {
	City, Provience, Country string
}

func main() {
	// address1 := Address{"Sleman" , "DIY" , "Idn"}
	// address2 := address1

	// address2.City = "Bantul"
	// fmt.Println(address1)
	// fmt.Println(address2)

	var address1 Address = Address{"Sleman" , "DIY" , "Idn"}
	var address2 *Address = &address1 // pointer
	var address3 *Address = &address1

	address2.City = "Bantul"
	
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}