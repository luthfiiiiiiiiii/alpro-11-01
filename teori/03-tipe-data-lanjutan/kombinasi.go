package main

import "fmt"

func main() {
    var p, q int
    fmt.Scan(&p, &q)

    out1 := (p % 2 == 0) || (q % 2 == 0)
    out2 := (p % 2 != 0) && (q % 2 != 0)
    out3 := !(p == q)

    fmt.Println(out1, out2, out3)
}
