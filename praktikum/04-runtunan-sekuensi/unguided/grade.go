package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Masukkan nama siswa: ")
    namaInput, _ := reader.ReadString('\n')
    nama := strings.TrimSpace(namaInput)

    var nilai int
    fmt.Print("Masukkan nilai siswa: ")
    fmt.Scan(&nilai)

    var grade string
    switch {
    case nilai >= 90 && nilai <= 100:
        grade = "A"
    case nilai >= 80 && nilai < 90:
        grade = "B"
    case nilai >= 70 && nilai < 80:
        grade = "C"
    case nilai >= 60 && nilai < 70:
        grade = "D"
    default:
        grade = "F"
    }

    fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
