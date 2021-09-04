package main

import (
	"fmt"
)

//
func main() {
	var tahun int
	fmt.Print("Masukkan Tahun : ")
	fmt.Scanln(&tahun)
	if tahun%4 == 0 && tahun%100 != 0 || tahun%400 == 0 {
		fmt.Println("Tahun", tahun, "adalah tahun kabisat")
	} else {
		fmt.Println("Tahun", tahun, "bukan tahun kabisat")
	}
}
