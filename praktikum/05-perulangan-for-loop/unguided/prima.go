package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlahFaktor := 0

	fmt.Print("Faktor dari ", n, " adalah : ")

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}

	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("ya, bilangan tersebut adalah bilangan prima")
	} else {
		fmt.Println("Tidak, bilangan tersebut bukan bilangan prima")
	}
}