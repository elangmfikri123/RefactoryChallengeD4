package main

import "fmt"

func main() {
	bil := []int{1, 2, 3, 5, 7, 11, 12, 32, 45, 65, 90, 70, 80, 101}

	fmt.Println()
	fmt.Println("Daftar bilangan =", bil)

	fmt.Print("Bilangan Genap = ")
	for i := 0; i < len(bil); i++ {
		if bil[i]%2 == 0 {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Bilangan Ganjil = ")
	for i := 0; i < len(bil); i++ {
		if bil[i]%2 != 0 {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Kelipatan 5 = ")
	for i := 0; i < len(bil); i++ {
		if bil[i]%5 == 0 {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Bilangan Prima = ")
	for i := 0; i < len(bil); i++ {
		var isPrime bool = true

		for j := 2; j < i; j++ {
			if bil[i]%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Print(bil[i], " ")
		}
	}

	fmt.Println()
	fmt.Print("Bilangan Prima kurang dari 100 = ")
	for i := 0; i < len(bil); i++ {
		var isPrime bool = true

		for j := 2; j < i; j++ {
			if bil[i]%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime && bil[i] < 100 {
			fmt.Print(bil[i], " ")
		}
	}
	fmt.Println()
}
