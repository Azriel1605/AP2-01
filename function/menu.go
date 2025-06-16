package function

import (
	"fmt"
)

func MainMenu() {

	fmt.Println("+------------------------------+")
	fmt.Println("|             MENU             |")
	fmt.Println("+------------------------------+")
	fmt.Println("| 1. Modifikasi Aset           |") 
	fmt.Println("| 2. Transaksi                 |") 
	fmt.Println("| 3. Portofolio                |") 
	fmt.Println("| 4. Histori Pembelian         |") 
	fmt.Println("| 5. Cari Aset Crypto          |") 
	fmt.Println("| 6. Urutkan Aset (harga)      |") 
	fmt.Println("| 7. Urutkan Portofolio (harga)|") 
	fmt.Println("| 8. Tampilkan Semua Aset      |") 
	fmt.Println("| 9. Exit                      |")
	fmt.Println("+------------------------------+")
}

func ModifyCryptoMenu() {

	fmt.Println("\n+-----------------------------+")
	fmt.Println("|-_-_- Modifikasi Crypto -_-_-|")
	fmt.Println("+-----------------------------+")
	fmt.Println("| 1. Tambah Crypto            |")
	fmt.Println("| 2. Hapus Crypto             |")
	fmt.Println("| 3. Ubah Harga Crypto        |")
	fmt.Println("| 4. Kembali ke Menu Utama    |")
	fmt.Println("+-----------------------------+")
}

func TransaksiMenu() {
	
	fmt.Println("\n+----------------------------+")
	fmt.Println("| 1. Beli Crypto             |")
	fmt.Println("| 2. Jual Crypto             |")
	fmt.Println("| 3. Kembali ke Menu Utama   |")
	fmt.Println("+----------------------------+")

}

