package main

import "fmt"

func main() {

	var t int
	fmt.Print("input t(jumlah transaksi): ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0;

		if hari == "Senin" || "Selasa" || "Rabu" || "Kamis" {
			if jenisFilm = "Reguler" {
				hargaSatuan = 40000
			} else if jenisFilm == "3D" {
				hargaSatuan = 60000
			}else if jenisFilm == "IMAX"{
				hargaSatuan = 80000}
		}

		if hari == "Jumat" || "Sabtu" || "Minggu" {
			if jenisFilm = "Reguler" {
				hargaSatuan = 50000
			} else if jenisFilm == "3D" {
				hargaSatuan = 70000
			}else if jenisFilm == "IMAX"{
				hargaSatuan = 90000}
		}

		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10)/100
			totalHarga = totalHarga - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalHarga)

        totalPendapatan = totalPendapatan + totalHarga
	}
	fmt.Println("output total pendapatan:", totalPendapatan)
}