package main

import "fmt"

func main() {
	for {
		var modal, hargaJual float64
		var jumlah int
		var diskon, pajak float64

		fmt.Println("Kalkulator Keuntungan")
		fmt.Print("Masukkan modal per Unit: ")
		fmt.Scan(&modal)
		fmt.Print("Masukkan harga jual per Unit: ")
		fmt.Scan(&hargaJual)
		fmt.Print("Masukkan jumlah Unit terjual: ")
		fmt.Scan(&jumlah)
		fmt.Print("Diskon%: ")
		fmt.Scan(&diskon)
		fmt.Print("PPN/Pajak: ")
		fmt.Scan(&pajak)

		//Diskon+Pjk
		hargaSetelahDiskon := hargaJual - (hargaJual * diskon / 100)
		hargaFinal := hargaSetelahDiskon + (hargaSetelahDiskon * pajak / 100)

		//Keuntungan
		untungPerUnit := hargaFinal - modal
		totalUntung := untungPerUnit * float64(jumlah)

		fmt.Println("-------------------------------------------")
		fmt.Printf("harga Setelah Diskon		:  Rp %.2f\n", hargaSetelahDiskon)
		fmt.Printf("Harga Setelah pajak		:  Rp %.2f\n", hargaFinal)
		fmt.Printf("Untung per Unit  		:  Rp %.2f\n", untungPerUnit)
		fmt.Printf("Total Keuntungan  		:  Rp %.2f\n", totalUntung)
		fmt.Println("-------------------------------------------")

		var ulang string
		fmt.Printf("Hitung Lagi? (y/n): ")
		fmt.Scan(&ulang)

		if ulang != "y" && ulang != "Y" {
			fmt.Println("Program selesai.Thanks")
			break
		}
	}
}
