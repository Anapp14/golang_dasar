package main

import "fmt"

func main12 () {
	var nilaiAkhir = 80
	var absen = 80

	var luluNilaiAkhir = nilaiAkhir > 70
	var luluAbsen = absen >= 80

	var lulus = luluNilaiAkhir && luluAbsen
	fmt.Println("kelulusan: ", lulus)
}