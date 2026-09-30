package main

import "fmt"

//menginisialisasikan adalah langsung memasukan isi(value) ke dalam variable

func main() {
//isi(value) di dalam variable harus sesuai dengan tipe datanya

//nama variable dan tipe datanya
	var name string //deklarasi

//isi(value) di dalam variable
	name = "luthfi shibghotillah"
	fmt.Println("halo nama saya", name)

//variable bisa menggunakan kata kunci :=
	namaAkhir := "shibghotillah"
	fmt.Println("nama akhir saya adalah", namaAkhir)

//bisa membuat lebih dari 1 nama variable dalam 1 variable
var (
	namaLengkap = "luthfi shibghotillah"
	namaDepan = "luthfi"
)//inisialisasi

fmt.Println(namaLengkap)
fmt.Println(namaDepan)

}