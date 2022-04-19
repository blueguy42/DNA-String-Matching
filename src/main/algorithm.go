package main

import (
	"fmt"
)

func kmpBorderFunc(pola string) []int {
	// Border function dari KMP merupakan array yang berisi panjang prefix maksimal yang sama dengan sufix untuk tiap k nya

	// KAMUS
	i := 0 // Menunjuk ke huruf terakhir pada prefix yang cocok dengan sufix untuk suatu k
	runePola := []rune(pola) // Konversi string menjadi array of rune agar bisa mengakses tiap char
	panjangPola := len(runePola)
	var border = make([]int,panjangPola-1) // Menggunakan slice karena array pada golang tidak bisa ngefix sizenya dari variabel
	
	// ALGORITMA
	for k := 1 ; k < len(border); k++ {
		for (runePola[k] != runePola[i] && i != 0) { // Pencarian i terlebih dahulu apabila huruf terakhirnya tidak sama
			i = border[i-1] // Dengan cara memperkecil i hingga ditemukan pasangan prefix-sufix yang cocok atau hingga habis tidak ditemukan
		}
		if runePola[k] == runePola[i]{ // Kasus pencarian i menemukan pasangan cocok, maka border function di k adalah panjang dari prefix
			i += 1
			border[k] = i
		} else { // Kasus pencarian i tidak menemukan pasangan cocok, sehingga border functionnya adalah 0
			border[k] = 0
		}
	}

	return border
}

func kmp(dna string, pola string) bool{
	// Memeriksa apakah dalam suatu dna terdapat pola atau tidak menggunakan KMP, return true untuk iya dan false untuk tidak

	// KAMUS
	i := 0 // merujuk ke huruf pada DNA yang sedang diperiksa
	j := 0 // merujuk ke huruf pada pola yang sedang diperiksa
	runeDNA := []rune(dna) // Konversi string menjadi array of rune agar bisa mengakses tiap char
	runePola := []rune(pola) 
	panjangDNA := len(runeDNA) 
	panjangPola := len(runePola)
	found := false
	borderFunction := kmpBorderFunc(pola)

	// ALGORITMA
	for (i < panjangDNA && !found){
		if (runeDNA[i] == runePola[j]){ // Kasus belum found, tapi huruf yang diperiksa cocok maka cek next huruf
			i += 1
			j += 1
		} else { // Kasus hurufnya tidak cocok
			if (j == 0) {
				i += 1 // Kasus kesalahan ditemukan sejak huruf pertama, maka geser i nya
			} else {
				j = borderFunction[j-1] // kasus kesalahan ditemukan selain huruf pertama, majukan j sesuai border function
			}
		}

		if (j == panjangPola){ // Kasus sebanyak panjang pola huruf sudah diperiksa dan betul, maka ketemu
			found = true
		}

	}

	return found
}

func booyermoreLastIndex(dna string, pola string) map[string]int {
	// KAMUS
	last := make(map[string]int) // Menggunakan map untuk menunjukkan lokasi huruf terakhirnya
	runePola := []rune(pola) // Konversi string menjadi array of char agar bisa mengakses tiap char
	runeDNA := []rune(dna)
	panjangPola := len(runePola)
	panjangDNA := len(dna)

	// ALGORITMA
	for i := 0; i < panjangDNA ; i++ { // Menginisiasi tiap huruf dalam DNA sebagai -1 terlebih dahulu
		last[string(runeDNA[i])] = -1
	}

	for i := 0; i < panjangPola ; i++ { //Mengisi index terakhir sebuah char ditemukan dalam pola
		last[string(runePola[i])] = i
	}

	return last
}

func booyermore(dna string, pola string) bool{
	// KAMUS

	// ALGORITMA
	return false;
}

//func main() {
	// TES SINTAKS DASAR
	// kalimat := "IPnya kiky adalah "
	// IPK := 4
	// fmt.Print(kalimat, ipk)

	// TES BORDER FUNCTION KMP
	// pola1 := "ababababba"
	// var border []int
	// border = kmpBorderFunc(pola1)
	// for i := 0; i< len(border); i++{
	// 	fmt.Print(border[i], " ")
	// }

	// TES KMP 
	// fmt.Print(kmp("abcabsda","bsda"))

	// TEST BOOYERMOORE LAST INDEX FUNCTION
	// last := booyermoreLastIndex("abcd","abacab")
	// fmt.Print(last)

//}