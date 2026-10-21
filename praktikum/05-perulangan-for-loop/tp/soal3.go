package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)

    sum := 0
    fmt.Print("Bilangan ganjil dari 1 hingga ", n, ": ")

    for i := 1; i <= n; i++ {
        if i%2 != 0 {
            fmt.Print(i, " ")
            sum += i
        }
    }

    fmt.Println("\nJumlah =", sum)
}
