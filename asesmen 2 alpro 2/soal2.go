package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, N int, nim string) int {
	max := -1

	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}

	return max
}

func main() {
	var T arrayMahasiswa
	var N int
	var nimCari string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("\nMasukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(T, N, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(T, N, nimCari)

	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, nilaiTerbesar)
	} else {
		fmt.Println("NIM tidak ditemukan")
	}
}
