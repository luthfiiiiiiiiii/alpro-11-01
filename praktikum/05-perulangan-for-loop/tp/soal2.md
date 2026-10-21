perbedaan for (3 komponen) dengan for (kondisi saja) adalah for (yang 3 komponen) ketiga komponennya diletakan setelah perintah perulangan, sedangkan for (kondisi saja) inisialisasi dan post statementnya diletakan diluar perintah

1. contoh for 3 komponen

package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        fmt.Println(i)
    }
}

2. contoh for kondisi

package main

import "fmt"

func main() {
    j := 1
    for j <= 3 {
        fmt.Println(j)
        j++
    }
}
