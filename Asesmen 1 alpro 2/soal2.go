package main

import "fmt"

func biayaDomestikPerHari() float64 {
	return 70000 + 250000 + 300000
}

func biayaMancanegaraPerHari() float64 {

	return biayaDomestikPerHari() * 1.5
}

func hitungHariTanggungan(lama int, tujuan string) int {
	if tujuan == "domestik" {
		if lama > 3 {
			return 3
		}
		return lama
	} else {
		if lama > 8 {
			return 8
		}
		return lama
	}
}

func hitungBiayaTotal(jumlah, lama int, tujuan string) float64 {
	var biayaPerHari float64
	var hariTanggungan int

	if tujuan == "domestik" {
		biayaPerHari = biayaDomestikPerHari()
	} else {
		biayaPerHari = biayaMancanegaraPerHari()
	}
	hariTanggungan = hitungHariTanggungan(lama, tujuan)
	return float64(jumlah) * biayaPerHari * float64(hariTanggungan)
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	biaya = hitungBiayaTotal(jumlah, lama, tujuan)
	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
