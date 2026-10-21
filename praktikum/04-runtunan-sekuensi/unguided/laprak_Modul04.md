# <h1 align="center">Tugas Pendahuluan Modul 04</h1>
<p align="center">Luthfi Shibghotillah - 109092600003</p>

### 1. soal tp 1

```go
package main

import "fmt"

func main() {

	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if (intOther + 2 * intNum != 30 || !(sngNum > 0) ) {
		fmt.Println("Beep")	
	} else {
		fmt.Println("false")
	}

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/Screenshot 2026-10-14 095302.png)


#### Deskripsi
menggunakan operator or(||) untuk menentukan salah satunya benar

## Kesimpulan
ada banyak operator yang bisa digunakan untuk menentukan true atau false

### 2. soal tp 2

```go
package main

import "fmt"

func main() {
    x := 10
    y := 5
    z := 15
    result := 0

    if x > 5 {
        if y < 10 {
            result = x + y
        } else {
            result = x - y
        }
    }

    if z > 10 && x == 10 {
        result += z
    } else {
        result = z - x
    }

    if x == 10 || y > 10 {
        result += 5
    } else if y == 5 && z > 10 {
        result -= 5
    } else {
        result *= 2
    }

    if !(x < 15 && y < 10) {
        result += 10
    } else {
        result -= 10
    }

    fmt.Println(result)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/)


#### Deskripsi
menggunakan perintah if dan else untuk menentukan kondisi tertentu

## Kesimpulan
kesimpulannya adalah kita bisa menentukan kondisi seperti apa yang kita inginkan untuk menjalankan sebuah program

### 3. soal tp 3

```go
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

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/)


#### Deskripsi
penggunaan perintah switch case

## Kesimpulan
penggunaan switch case sedikit mirip dengan if dan else

### 4. soal tp 4

```go
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

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/)


#### Deskripsi
penggunaan perintah switch case

## Kesimpulan
penggunaan switch case sedikit mirip dengan if dan else