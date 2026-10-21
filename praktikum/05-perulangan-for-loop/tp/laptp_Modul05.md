# <h1 align="center">Tugas Pendahuluan Modul 03 - modulo, boolean, & konversi</h1>
<p align="center">Luthfi Shibghotillah - 109092600003</p>

### 1. soal tp 1

```go

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided]()

#### Deskripsi
fungsi utama dari for adalah untuk menjalankan perintah perulangan
3 komponen for yaitu:
    1. inisialisasi : yaitu menetapkan nilai awal suatu variabel
    2. kondisi : memeriksa apakah perintah perulangan masih harus dijalankan atau dihentikan
    3. post statement : mengubah nilai variabel berdasarkan perintahnya

## Kesimpulan

### 2. soal tp 2

```go

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided]()


#### Deskripsi
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

## Kesimpulan

### 3. soal tp 3

```go
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

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/fototp/fotosoal3.png)


#### Deskripsi
program ini diawali dengan memasukan bilangan bulat psitif, lalu program akan menampilkan bilangan ganjil yang tersedia dari 1 sampai bilangan yang kita masukan, lalu program akan menjumlahkan bilangan ganjil yang ditemukan

## Kesimpulan
kesimpulannya adalah kita bisa membuat program untuk menyeleksi sebuah bilangan

### 4. soal tp 4

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/fototp/fotosoal4.png)


#### Deskripsi
program ini akan menampilkan sebuah tabel perkalian dengan banyak tabel berdasarkan angka yang kita masukan

## Kesimpulan
kesimpulannya adalah kita bisa mengulangi sebuah perintah seanyak yang kita inginkan
