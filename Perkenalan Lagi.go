package main
import "fmt"

func main() {
	//VariabelDataUser
	var (
		Nama 			string 	= "Fatih"
		Umur		 	int 	= 17
		Pendidikan 		string 	= "SMK"
		Status 			bool 	= false
		hobi 			string 	= "Bermain Game"
	)
							//statusPernikahan
						statusMap := map[bool]string {
							true: "Sudah Menikah",
							false: "Belum Menikah",
						}
						
		//Printf(Output)
		fmt.Println("==== Profil User ===")	
		fmt.Printf("Nama				:	%s\n", Nama)
		fmt.Printf("Umur				:	%d\n", Umur)
		fmt.Printf("Status Pernikahan 		:	%s\n", statusMap[Status])
		fmt.Printf("Pendidikan			:	%s\n", Pendidikan)
		fmt.Printf("Hobi Saya 			:	%s\n", hobi)
}

//{
//HasilPembelajaran tanggal 8 bulan 12 2025

//A.fungsidasar:

//	A1.var(variabel) = Tempat atau lokasi untuk pernyimpanan tipedata.Diawali dengan var {tipedata} atau dengan :=
// 	A2.Tipedata = 
	//	1.string(data/tulisan/huruf.Menggunakan = "" untuk input).Contoh = var Nama = "User"
	//	2.int(angka.Hanya menggunakan = untuk input).Contoh = var int = 25
	//	3.bool = menentukan true/false.Contoh = var PunyaUang bool = true.
//	A3.Print(Printf,Println,Print)
//	A4.scan.Untuk input dari user ke kode program.Contoh = fmt.Scan(&var) atau fmt.Scanln(&var)
//}