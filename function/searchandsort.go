package function

import (
	"IF-48-10_AP2-01_TUBES-01/data"
	"fmt"
)

func LengtOfExistingData(A *data.TabCrypto) int {
	// I.S A terdefinisi
	// F.S Mengembalikan jumlah data yang terisi dalam A

	var i int
	for A[i].Name != "" && i < data.NMAX {
		i++
	}
	return i
}

func LengtOfExistingDataPorto(P *data.TabPorto) int {
	// I.S A terdefinisi
	// F.S Mengembalikan jumlah data yang terisi dalam P
	
	var i int
	for P[i].Name != "" && i < data.NMAX {
		i++
	}
	return i
}

func SequencialSearchCrypto(A *data.TabCrypto, name string, showNotFound bool) int {
	// I.S A terdefinisi, name terdefinisi, showNotFound terdefinisi
	// F.S Mengembalikan indeks dari nama crypto yang dicari, jika tidak ditemukan mengembalikan -1
	var i int

	for i = 0; i < data.NMAX; i++ {
		if A[i].Name == name {
			return i
		}
	}
	// Jika showNotFound bernilai true, tampilkan pesan bahwa crypto tidak ditemukan
	if showNotFound {
		fmt.Println("Crypto tidak ditemukan")
	}
	return -1
}


func SequencialSearchPorto(P *data.TabPorto, name string) int {
	// I.S P terdefinisi, name terdefinisi
	// F.S Mengembalikan indeks dari nama crypto yang dicari menggunakan algoritma sequential search, jika tidak ditemukan mengembalikan -1
	var i int
	for i = 0; i < data.NMAX; i++ {
		if P[i].Name == name {
			return i
		}
	}
	fmt.Println("Crypto tidak ditemukan")
	return -1
}

func BinarySearchCrypto(A data.TabCrypto, name string) int {
	// I.S A terdefinisi, name terdefinisi
	// F.S Mengembalikan indeks dari nama crypto yang dicari menggunakan algoritma binary search, jika tidak ditemukan mengembalikan -1
	var low, mid, high int
	low = 0
	high = LengtOfExistingData(&A) - 1
	for low <= high {
		mid = (low + high) / 2
		if A[mid].Name == name {
			return mid
		} else if A[mid].Name < name {
			low = mid + 1
		} else if A[mid].Name > name {
			high = mid - 1
		}
	}

	return -1 // Data not found
}

func SelectionSortCrypto(A *data.TabCrypto) {
	// I.S A terdefinisi
	// F.S A terurut berdasarkan harga crypto secara ascending menggunakan algoritma selection sort

	var i, j, n, min int
	var temp data.TCrypto
	n = LengtOfExistingData(A) //Menampilkan banyakya data yang terisi

	for i = 0; i < n; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if A[j].Price < A[i].Price {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}

}

func SelectionSortPorto(P *data.TabPorto) {
	// I.S P terdefinisi
	// F.S P terurut berdasarkan harga pasar crypto secara ascending menggunakan algoritma selection sort

	var i, j, n, min int
	var temp data.TPorto
	n = LengtOfExistingDataPorto(P) //Menampilkan banyakya data yang terisi

	for i = 0; i < n; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if P[j].MarketPrice < P[i].MarketPrice {
				min = j
			}
		}
		temp = P[i]
		P[i] = P[min]
		P[min] = temp
	}
}

func InsertionSortPorto(P *data.TabPorto) {
	// I.S P terdefinisi
	// F.S P terurut berdasarkan harga pasar crypto secara ascending menggunakan algoritma insertion sort

	var i, n, j int
	var temp data.TPorto

	n = LengtOfExistingDataPorto(P) //Menampilkan banyakya data yang terisi
	for i = 1; i < n; i++ {
		j = i - 1
		temp = P[i]
		for j >= 0 && P[j].MarketPrice > temp.MarketPrice {
			P[j+1] = P[j]
			j--
		}
		P[j+1] = temp
	}
}