# <h1 align="center">Tugas unguided Modul 03 - konversi</h1>
<p align="center">Luthfi Shibghotillah - 109092600003</p>

### 1. konversi hari

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi)


#### Deskripsi

## Kesimpulan

### 2. Konversi

```go
package main

import "fmt"

func main() {

    var c, r float64

    fmt.Scan(&c)

    r = (4.0 / 5.0) * c

    fmt.Printf("%.2f\n", r)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi)


#### Deskripsi
menggunakan perintah print dengan format "%.2f", sehingga akan menghasilkan bilangan real yang awalnya memiliki beberapa angka setelah koma(,) menjadi hanya 2 angka setelah koma

## Kesimpulan
kesimpulannya adalah kita bisa mengatur banyaknya angka yang ingin kita tampilkan dengan mnggunakan format tertentu