package main

import "fmt"

func main() {
    var tahun int
    var bulan string

    fmt.Scan(&tahun)
    fmt.Scan(&bulan)

    kabisat := (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)

    var hari int
    benar := true

    switch bulan {
    case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
        hari = 31
    case "Apr", "Jun", "Sep", "Nov":
        hari = 30
    case "Feb":
        if kabisat {
            hari = 29
        } else {
            hari = 28
        }
    default:
        benar = false
    }

    if benar {
        fmt.Printf("%s %d memiliki %d hari\n", bulan, tahun, hari)
    } else {
        fmt.Printf()
    }
}
