package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Masukkan bilangan yang dikalikan (a): ")
    fmt.Scan(&a)
    fmt.Print("Masukkan batas perkalian (b): ")
    fmt.Scan(&b)

    fmt.Println("Tabel perkalian", a, "dari", a, "x 1 sampai", a, "x", b)

    for i := 1; i <= b; i++ {
        fmt.Printf("%d x %d = %d\n", a, i, a*i)
    }
}
