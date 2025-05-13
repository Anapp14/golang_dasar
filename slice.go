package main

import "fmt"

func main14() {
	names := [...]string{"Anap", "Budi", "Caca", "Dede", "Eka", "Fani", "Gina", "Hana", "Ika", "Joko"}

	slice1 := names[4:8]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	var slice4 []string = names[2:5]
	fmt.Println(slice4)

	// Mengubah isi slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]

	// Mengubah isi slice dengan append
	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "Ups!"
	fmt.Println(daySlice2) // [Ups! Minggu baru Libur baru]
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]

	// Menggunakan make untuk membuat slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "anap"
	newSlice[1] = "anap"
	//newSlice[2] = "anap" Error: index out of range, harus menggunakan append

	fmt.Println(newSlice) // [anap anap]
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	newSlice2 := append(newSlice, "anap")
	fmt.Println(newSlice2) // [anap anap anap]
	fmt.Println(len(newSlice2)) // 3
	fmt.Println(cap(newSlice2)) // 5

	newSlice2[0] = "budi"
	fmt.Println(newSlice2) // [budi anap anap]
	fmt.Println(newSlice)

	// Menggunakan copy untuk menyalin slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice) // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]
	fmt.Println(toSlice) // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]

	// Array vs Slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray) // [1 2 3 4 5]
	fmt.Println(iniSlice) // [1 2 3 4 5]
}