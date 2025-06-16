package function // DONE

import (
	"IF-48-10_AP2-01_TUBES-01/data"
	"IF-48-10_AP2-01_TUBES-01/v"
	"fmt"
)

func AddPortofolio(A *data.TabCrypto,
	P *data.TabPorto,
	name string,
	quantity int) {
	var i, j, idxC int
	i = SequencialSearchPorto(P, name, false)
	if i != -1 {
		P[i].Quantity += quantity

	} else {
		idxC = SequencialSearchCrypto(A, name, false)
		for P[j].Name != "" && j < data.NMAX {
			j++
		}
		P[j].Name = name
		P[j].Quantity = quantity
		P[j].BuyPrice = A[idxC].Price
		P[j].MarketPrice = A[idxC].Price
	}
}

func SellPortofolio(P *data.TabPorto, name string, n int){
	// I.S P terdefinisi, name terdefinisi, n terdefinisi
	// F.S Menjual crypto dari portofolio P
	// Jika crypto tidak ditemukan, tampilkan pesan "Tidak ditemukan"

	var i int = SequencialSearchPorto(P, name, true)
		
	if i != -1 { // Kalau ada nama di portofolio
		if P[i].Quantity < n {
				fmt.Println("Lot lebih Banyak dari yang ada di portofolio")
		} else {
			P[i].Quantity -= n
			if P[i].Quantity == 0 {
				P[i].Name = ""
				P[i].Quantity = 0
				P[i].BuyPrice = 0
				P[i].MarketPrice = 0
			}
			fmt.Println("Crypto dalam portofolio berhasil dijual")
		}
	}
}

func RmPortofolio(P *data.TabPorto,
	name string) {

	// I.S P terdefinisi, name terdefinisi
	// F.S Menghapus crypto dari portofolio P

	var i int = SequencialSearchPorto(P, name, true)
	if i != -1 {
		P[i].Name = ""
		P[i].Quantity = 0
		P[i].BuyPrice = 0
		P[i].MarketPrice = 0
	}
}

func ShowPortofolio(P *data.TabPorto) {
	// I.S P terdefinisi
	// F.S Menampilkan semua crypto dalam portofolio P

	var i int
	var equity int

	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("|                         PORTOFOLIO                         |")
	fmt.Println("+------------------------------------------------------------+")
	fmt.Println("|       Crypto        |   Q   |  Harga Beli  |   Harga Jual  |")
	fmt.Println("+------------------------------------------------------------+")
	for P[i].Name != "" && i < data.NMAX {
		fmt.Printf("| %-19s | %-5d | %-12.0f | %-13.0f |\n", P[i].Name, P[i].Quantity, P[i].BuyPrice, P[i].MarketPrice)
		equity += P[i].Quantity * int(P[i].MarketPrice)
		i++
	}
	equity += int(v.Saldo)
	fmt.Println("+------------------------------------------------------------+")
	fmt.Printf("| %-19s : %36.0f |\n", "Sisa Modal", float64(v.Saldo))
	fmt.Println("+------------------------------------------------------------+")
	fmt.Printf("| %-19s : %36.0f |\n", "Total Ekuitas", float64(equity))
	fmt.Println("+------------------------------------------------------------+")

}
