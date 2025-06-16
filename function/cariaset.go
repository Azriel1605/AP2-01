package function

import (
	"IF-48-10_AP2-01_TUBES-01/data"
	"fmt"
)

func CariCrypto(A *data.TabCrypto) {
	// I.S A terdefinisi
	// F.S Mencari crypto berdasarkan nama menggunakan binary search

	var i int
	var name string
	fmt.Println("Masukkan nama crypto yang ingin dicari: ")
	fmt.Scan(&name)
	i = BinarySearchCrypto(*A, name)
	if i != -1 {
		fmt.Printf("\n\nLokasi Index : %d\nNama Crypto: %s\nHarga Beli: %.0f\n", i, A[i].Name, A[i].Price)
	} else {
		fmt.Println("Crypto tidak ditemukan")
	}
}