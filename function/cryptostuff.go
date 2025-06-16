package function

import (
	"IF-48-10_AP2-01_TUBES-01/data"
	"IF-48-10_AP2-01_TUBES-01/v"
	"fmt"
)

func AddCrypto(A *data.TabCrypto) {
	// I.S A terdefinisi
	// F.S Menambahkan crypto baru ke dalam array A

	var i, j, n, k int
	var name string
	var price float64

	fmt.Println("Masukkkan jumlah cyrpto yang akan ditambahkan")
	fmt.Scan(&n)

	// Bikin looping
	for k=0; k < n; k++ {
		fmt.Println("Masukkan nama crypto dan harga: ")
		fmt.Scan(&name, &price)
		
		i = SequencialSearchCrypto(A, name, false)
		
		if i == -1 {
			for A[j].Name != "" && j < data.NMAX {
				j++
			}
			if j >= data.NMAX {
				fmt.Println("Array Crypto sudah mencapai batas maksimal")
			} else {
				A[j].Name = name
				A[j].Price = price
			}
		} else {
			fmt.Println(name, " sudah terdaftar")
		}
	}


}

func DeleteCrypto(A *data.TabCrypto) { // DONE
	// I.S A terdefinisi
	// F.S Menghapus crypto dari array A

	var i, j int
	var name string
	fmt.Println("Masukkan nama crypto yang ingin dihapus: ")
	fmt.Scan(&name)
	i = SequencialSearchCrypto(A, name, true)
	if i != -1 {
		for j = i; j < data.NMAX-1; j++ {
			A[j] = A[j+1]
		}
		A[data.NMAX-1].Name = ""
		A[data.NMAX-1].Price = 0
		fmt.Println("Crypto berhasil dihapus")
	}
}

func UpdateCrypto(A *data.TabCrypto) { // DONE
	// I.S A terdefinisi
	// F.S Mengupdate harga crypto dalam array A

	var i, idxP int
	var name string
	fmt.Println("Masukkan nama crypto yang ingin diubah: ")
	fmt.Scan(&name)
	i = SequencialSearchCrypto(A, name, true)
	if i != -1 {
		fmt.Println("Masukkan harga crypto yang baru: ")
		fmt.Scan(&A[i].Price)
		idxP = SequencialSearchPorto(&v.Porto, name)
		if idxP != -1 {
			v.Porto[idxP].MarketPrice = A[i].Price
		}
	}
}

func ShowCrypto(A *data.TabCrypto) {
	// I.S A terdefinisi
	// F.S Menampilkan semua crypto dalam array A
	
	var i int
	fmt.Println("+--------------------------------------+")
	fmt.Println("|            CRYPTOCURRENCY            |")
	fmt.Println("+--------------------------------------+")
	fmt.Println("| NO |       Crypto        |   Harga   |")
	fmt.Println("+--------------------------------------+")
	for i = 0; i < data.NMAX; i++ {
		if A[i].Name != "" {
			fmt.Printf("| %-2d | %-19s | %-9.0f |\n", i+1, A[i].Name, A[i].Price)
		}
	}
	fmt.Println("\n+--------------------------------------+")
}