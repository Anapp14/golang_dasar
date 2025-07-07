package main

import "fmt"

type User struct {
	Username    string
	IsLoggedIn  bool
}

// TODO: Lengkapi method ini
func (u *User) Login() {
	
}

func main() {
	user := User{"akhnaf", false}
	fmt.Println("Sebelum login:", user.IsLoggedIn) // false

	user.Login()

	fmt.Println("Setelah login:", user.IsLoggedIn) // true
}
