package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlahLolos := 0

	for i := 1; i <= n; i++ {
		var PI, V, GA, FS, CI, TC int

		fmt.Scan(&PI, &V, &GA, &FS, &CI, &TC)

		total := PI + V + GA + FS + CI + TC
		rataRata := float64(total) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Ya")
			jumlahLolos++
		} else {
			fmt.Println("Tidak")
		}
	}

	fmt.Println(jumlahLolos)
}