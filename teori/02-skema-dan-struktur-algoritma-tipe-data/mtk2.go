package main

import ("fmt")

func main() {
    var x float64

    fmt.Scan(&x)

    f := (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
    fmt.Println("Hasilnya adalah", f)
	
}
