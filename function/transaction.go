package function

import (
	"IF-48-10_AP2-01_TUBES-01/data"
	"IF-48-10_AP2-01_TUBES-01/v"
	"fmt"
)

func Beli(A *data.TabCrypto) {
	// I.S A terdefinisi
	// F.S Melakukan transaksi pembelian crypto, jika berhasil maka saldo virtual akan berkurang sesuai dengan harga crypto yang dibeli

	var i int
	var name string
	var quantity int

	fmt.Println("Masukkan nama crypto yang ingin dibeli: ")
	fmt.Scan(&name)

	i = SequencialSearchCrypto(A, name, false) // Jika crypto tidak ditemukan, i akan bernilai -1

	if i != -1 { // Jika crypto ditemukan, if akan dieksekusi
		fmt.Println("Masukkan jumlah crypto yang ingin dibeli: ")
		fmt.Scan(&quantity)
		if v.Saldo >= (A[i].Price * float64(quantity)) { // Periksa apakah saldo cukup untuk membeli crypto
			v.Saldo -= (A[i].Price * float64(quantity))

			// Masukkan transaksi ke hisotory
			HistoryRecord(&v.History, &v.IdxHistory, name, A[i].Price, quantity, "buy")
			AddPortofolio(A, &v.Porto, name, quantity)
			fmt.Println("Transaksi berhasil")
		} else {
			fmt.Println("Saldo tidak cukup")
		}
	}
}

// 
func Jual(P *data.TabPorto) {
	// I.S P terdefinisi
	// F.S Melakukan transaksi penjualan crypto, jika berhasil maka saldo virtual akan bertambah sesuai dengan harga crypto yang dijual

	var i int
	var name string
	var quantity int
	fmt.Println("Masukkan nama crypto yang ingin dijual: ")
	fmt.Scan(&name)
	i = SequencialSearchPorto(P, name, true)
	if i != -1 {
		fmt.Println("Masukkan jumlah crypto yang ingin dijual: ")
		fmt.Scan(&quantity)
		if P[i].Quantity >= quantity {
			v.Saldo += (P[i].MarketPrice * float64(quantity))
			// Masukkan transaksi ke hisotory
			HistoryRecord(&v.History, &v.IdxHistory, name, P[i].MarketPrice, quantity, "sell")
			SellPortofolio(P, name, quantity)
			fmt.Println("Transaksi berhasil")
		} else {
			fmt.Println("Jumlah crypto yang ingin dijual melebihi jumlah crypto yang dimiliki")
		}
	}
}
