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
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/foto%20tp/soaltp1.png)


#### Deskripsi
operator logika pada go lang digunakan untuk melakukan pengecekan satu atau lebih kondisi. pada kasus ini operator logika yang digunakan adalah OR(||), yang akan menghasikan nilai true jika salah satu saja kondisi yang bernilai true

## Kesimpulan
kesimpulannya adalah operator logika dapat membuat program melakukan pengecekan terhadap beberapa kondisi

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
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/foto%20tp/soaltp2.png)


#### Deskripsi
menggunakan perintah if dan else untuk membuat program yang akan menjalankan perintah sesuai dengan kondisi yang diberikan

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
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/foto%20tp/soaltp3&4.png)


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
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/foto%20tp/soaltp3&4.png)


#### Deskripsi
penggunaan perintah switch case

## Kesimpulan
penggunaan switch case sedikit mirip dengan if dan else