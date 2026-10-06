package main

import "fmt"

func main() {
	var x int

	fmt.Print("masukan nominal yang akan dipecah : ")
	fmt.Scan(&x)

	var sepuluhribuan int = x / 10000
	var sisa int = x % 10000

	var limaribuan int = x / 5000
	sisa = sisa % 5000

	var seribuan int = sisa / 1000

	fmt.Println(sepuluhribuan, limaribuan, seribuan)
}