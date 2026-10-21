package main

import (
    "fmt"
)

func main() {
    var penghasilan float64
    fmt.Print("Masukkan penghasilan (juta): ")
    fmt.Scan(&penghasilan)

    var pajak float64

    if penghasilan <= 50 {
        pajak = penghasilan * 0.05
    } else if penghasilan <= 100 {
        pajak = (50 * 0.05) + ((penghasilan - 50) * 0.10)
    } else if penghasilan <= 200 {
        pajak = (50 * 0.05) + (50 * 0.10) + ((penghasilan - 100) * 0.15)
    } else {
        pajak = (50 * 0.05) + (50 * 0.10) + (100 * 0.15) + ((penghasilan - 200) * 0.20)
    }

    fmt.Printf("Total pajak yang harus dibayar: %.2f juta\n", pajak)
}
