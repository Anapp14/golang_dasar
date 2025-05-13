package main

import "fmt"

func main10() {
	var a = 10
	var b = 20
	var c = 30

	var d = a + b + c

	fmt.Println("Hasil Penjumlahan", d)

	var e = 10
	var f = 20
	var g = 30

	var h = e * f * g

	fmt.Println("Hasil Perkalian", h)

	var i = 10
	var j = 20
	var k = 30

	var l = i - j - k

	fmt.Println("Hasil Pengurangan", l)

	var m = 18
	var n = 5

	var p = m / n

	fmt.Println("Hasil Pembagian", p)

	var z = 10
	z += 10 // z = z + 10
	fmt.Println("Hasil Penjumlahan", z)

	var u = 1
	u++ // u = u+1
	fmt.Println(u)
	u++
	fmt.Println(u)
	u--
	fmt.Println(u)
}