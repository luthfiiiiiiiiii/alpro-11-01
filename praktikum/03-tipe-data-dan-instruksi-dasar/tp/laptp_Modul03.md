# <h1 align="center">Tugas Pendahuluan Modul 03 - modulo, boolean, & konversi</h1>
<p align="center">Luthfi Shibghotillah - 109092600003</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {

	var y, x int

	fmt.Scan(&y, &x)

	sisa := y % x

	fmt.Println(sisa)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa/Screenshot%202026-09-30%20165158.png)


#### Deskripsi
menggunakan operator modulo(%) untuk mencari sisa hasil pembagian

## Kesimpulan
kita bisa menggunakan operator untuk membuat sebuah perhitungan

### 2. Konversi

```go
package main

import "fmt"

func main() {
	
	var mil float64

	fmt.Scan(&mil)

	km := mil * 1.6

	fmt.Printf("%.1f", km)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/Screenshot%202026-09-30%20165616.png)


#### Deskripsi
menggunakan perintah print dengan format "%.1f", sehingga akan menghasilkan bilangan real yang awalnya memiliki beberapa angka setelah koma(,) menjadi hanya 1 angka setelah koma

## Kesimpulan
kesimpulannya adalah kita bisa mengatur banyaknya angka yang ingin kita tampilkan dengan mnggunakan format tertentu

### 3. boolean

```go
package main

import "fmt"

func main() {

	var benar bool = true

	fmt.Scan(&benar)

	fmt.Println(benar)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/Screenshot%202026-09-30%20165616.png)


#### Deskripsi
menggunakan tipe data bool(boolean) sehingga kita memiliki salah satu kondisi yaitu true atau false

## Kesimpulan
penggunaan tipe data boolean hanya memiliki 2 kondisi yaitu true atau false