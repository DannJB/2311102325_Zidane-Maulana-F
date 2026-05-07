package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxTercepat := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxTercepat] {
			idxTercepat = i
		}
	}
	return idxTercepat
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (1 + tumbuh[i])
			fmt.Printf("%s %.0f\n", prov[i], prediksi)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	InputData(&prov, &pop, &tumbuh)

	var namaCari string
	fmt.Scan(&namaCari)

	idxTercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxTercepat])

	idxDicari := IndeksProvinsi(prov, namaCari)
	fmt.Printf("Data provinsi yang dicari : %s\n", prov[idxDicari])

	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	Prediksi(prov, pop, tumbuh)
}
