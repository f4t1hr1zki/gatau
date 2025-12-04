package main
import (
	"fmt"
	"time"
)
func main()  {
	var detik int
	fmt.Print("Masukkan Waktu(detik): ")
	fmt.Scan(&detik)
	
	fmt.Println("Timer dimulai...")
	time.Sleep(time.Duration(detik) * time.Second)
	
	fmt.Println("Waktu Habis")
}