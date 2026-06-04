package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func insertionSortDesc(t *tabPartai, n int) {
	for i := 1; i < n; i++ {
		key := t[i]
		j := i - 1
		for j >= 0 && t[j].suara < key.suara {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

func main() {
	var p tabPartai
	var n int = 0
	var suara int

	fmt.Println("Masukkan proses input suara :")
	fmt.Scan(&suara)

	for suara != -1 && n < NMAX {
		idx := posisi(p, n, suara)
		if idx == -1 {
			p[n].nama = suara
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
		fmt.Scan(&suara)
	}

	fmt.Println("Hasil Perhitungan suara :")

	if n == 0 {
		fmt.Println()
		return
	}

	insertionSortDesc(&p, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
	}
	fmt.Println()
}
