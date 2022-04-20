package main

// import (
// 	"fmt"
// )

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

func boyermooreLastIndex(dna string, pola string) map[string]int {
	// Mengembalikan sebuah map untuk tiap huruf pada pola, pada indeks terakhir keberapakah mereka ditemukan. Untuk huruf pada dna yang tidak ada pada pola dicatat -1

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

func boyermoore(dna string, pola string) bool{
	// Memeriksa apakah dalam suatu dna terdapat pola atau tidak menggunakan boyermoore, return true untuk iya dan false untuk tidak

	// KAMUS
	runePola := []rune(pola) // Konversi string menjadi array of char agar bisa mengakses tiap char
	runeDNA := []rune(dna)
	panjangPola := len(runePola)
	panjangDNA := len(runeDNA)
	i := panjangPola -1 // Menunjuk kepada huruf pada dna yang sedang diperiksa
	j := panjangPola -1 // menunjuk kepada huruf pada pola yang sedang diperiksa
	lastIndex := boyermooreLastIndex(dna,pola)
	found := false

	// ALGORITMA
	if (i < panjangDNA){ // Pencarian hanya dilakukan apabila i lebih kecil daripada panjang pola, karena jika pola lebih panjang daripada dna tidak dapat dicari
		for (i < panjangDNA && !found){
			if (runeDNA[i] == runePola[j]){ // Apabila hurufnya sama, maka periksa huruf sebelumnya. Teknik looking glass
				i -= 1
				j -= 1
			} else { // Kasus hurufnya tidak sama, maka kita menggeser i dan j agar align
				if (lastIndex[string(runeDNA[i])] != -1){
					if (lastIndex[string(runeDNA[i])] + 1 < j){ // CASE 1 : Huruf yang salah ada pada lastIndex, dan bisa digeser sampe di-align
						i += panjangPola - lastIndex[string(runeDNA[i])] - 1
					}else { // CASE 2 : Huruf yang salah ada pada lastIndex, namun tidak bisa digeser untuk di-align
						i += panjangPola - j
					}
				} else { // CASE 3 : Huruf yang salah pada dna tidak ditemukan pada lastIndex, geser sejauh panjangPola
					i += panjangPola
				}

				j = panjangPola - 1 // untuk me-reset j kembali ke huruf terakhir pada pola
			}

			if (j == -1){ // apabila semua huruf pola sudah diperiksa hingga j nya -1, maka sudah ditemukan
				found = true
			}
		}
	}

	return found;
}

// func main() {
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
	// fmt.Print(kmp("abcabsda","ashaoioahsoithhtash"))
	// fmt.Print(kmp("agctcgatcgatcgatc", "gatc"))

	// TEST BOYERMOORE LAST INDEX FUNCTION
	// last := boyermooreLastIndex("abcd","abacab")
	// fmt.Print(last)

	// TES boyermoore
	// fmt.Print(boyermoore("abcabsda","ashaoioahsoithhtash"))
	// fmt.Print(boyermoore("agctcgatcgatcgatc", "gatc"))
	// fmt.Print(boyermoore("abcdefgh", "fgh"))

// }