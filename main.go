package main

import (
	"IF-48-10_AP2-01_TUBES-01/function"
	"IF-48-10_AP2-01_TUBES-01/v"
	"fmt"

)

func main() {
	var choice int

	fmt.Println("Masukkan saldo virtual yang anda miliki: ")
	fmt.Scan(&v.Saldo)
	for choice != 9 {

		function.MainMenu()
		fmt.Scan(&choice)

		function.ClearScreen()
		switch choice {
		case 1:
			function.ModifyCryptoMenu()
			fmt.Scan(&choice)
			
			switch choice {
			case 1:
				function.AddCrypto(&v.Crypto)
			case 2:
				function.DeleteCrypto(&v.Crypto)
			case 3:
				function.UpdateCrypto(&v.Crypto)
			case 4:
				fmt.Println("Kembali ke Menu Utama")
			default:
				fmt.Println("Pilihan tidak valid")
			}

		case 2:

			function.TransaksiMenu()
			fmt.Scan(&choice)
			

			switch choice {
			case 1:
				function.Beli(&v.Crypto)
			case 2:
				function.Jual(&v.Porto)
			case 3:
				fmt.Println("Kembali ke Menu Utama")
			default:
				fmt.Println("Pilihan tidak valid")
			}

		case 3:
			function.ShowPortofolio(&v.Porto)
		case 4:
			function.ShowHistory(&v.History)
		case 5:
			function.CariCrypto(&v.Crypto)
		case 6:
			fmt.Println("Sebelum diurutkan")
			function.ShowCrypto(&v.Crypto)
			function.SelectionSortCrypto(&v.Crypto)
			fmt.Print("\n\n\n")
			fmt.Println("Setelah diurutkan")
			function.ShowCrypto(&v.Crypto)

		case 7:
			
			fmt.Println("Sebelum diurutkan")
			function.ShowPortofolio(&v.Porto)
			function.SelectionSortPorto(&v.Porto)
			fmt.Print("\n\n\n")
			fmt.Println("Setelah diurutkan")
			function.ShowPortofolio(&v.Porto)

		case 8:
			
			function.ShowCrypto(&v.Crypto)
		case 9:
			
			fmt.Println("Sampai Jumpa")

		}
	}
}
