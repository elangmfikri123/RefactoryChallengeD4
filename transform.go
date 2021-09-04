package main

import "fmt"

func main() {

	type murid struct {
		nama, kelas string
	}

	var s1 murid
	s1.nama = "Elang Muhammad Fikhri"
	s1.kelas = "3 SMK"

	fmt.Println("Nama :", s1.nama)
	fmt.Println("Kelas :", s1.kelas)
}
