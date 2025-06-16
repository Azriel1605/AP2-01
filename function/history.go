package function

import (
	"IF-48-10_AP2-01_TUBES-01/data"
	"fmt"
)

func HistoryRecord(H *data.TabHistory, idx *int, name string, price float64, quantity int, transactionType string) {
	// I.S Semua parameter terdefinisi
	// F.S Mencatat history transaksi ke dalam array H

	H[*idx].Name = name
	H[*idx].BuyPrice = price
	H[*idx].Quantity = quantity
	H[*idx].TotalPrice = price * float64(quantity)
	H[*idx].TransactionType = transactionType
	*idx++
}

func ShowHistory(H *data.TabHistory) {
	// I.S H terdefinisi
	// F.S Menampilkan semua history transaksi dalam array H

	var i int
	fmt.Println("+---------------------------------------------------------------------------+")
	fmt.Println("|                             History Transaksi                             |")
	fmt.Println("+---------------------------------------------------------------------------+")
	fmt.Println("|Name           |Buy Price    |Quantity |Total Transaction |Transaction Type|")
	for H[i].Name != "" && i < data.NMAX {
		fmt.Printf("|%-14s |%-12.0f |%-8d |%-17.0f |%-15s |\n",
			H[i].Name, H[i].BuyPrice, H[i].Quantity, H[i].TotalPrice, H[i].TransactionType)
		i++
	}
	fmt.Println("+---------------------------------------------------------------------------+")
}
