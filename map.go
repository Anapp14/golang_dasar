package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Anap"
	//person["age"] = "19"
	//person["address"] = "Yogyakarta"
	//person["hobby"] = "Coding"

	person := map[string]string{
		"name":    "Anap",
		"age":     "19",
		"address": "Yogyakarta",
		"hobby":   "Coding",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])
	fmt.Println(person["hobby"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Anap"
	book["publisher"] = "Anap Publisher"
	book["year"] = "2023"
	book["price"] = "100000"

	fmt.Println(book)
	delete(book, "price")
	fmt.Println(book)
}