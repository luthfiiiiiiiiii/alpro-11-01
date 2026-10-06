package main

import "fmt"

func main() {
    var d, y, m, w, sisanya int

    fmt.Scan(&d)

    y = d / (12 * 30)
    sisa := d % (12 * 30)

    m = sisa / 30
    sisa = sisa % 30

    w = sisa / 7
    sisanya = sisa % 7

    fmt.Println(y, m, w, sisanya)

}
