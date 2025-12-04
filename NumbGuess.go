package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
		rand.Seed(time.Now().UnixNano())
		secret := rand.Intn(100) + 1
		var guess int
			fmt.Println("Tebak Angka antara 1-100!")
		
		for{
			fmt.Println("Tebakanmu: ")
			fmt.Scan(&guess)
			
				if guess < secret{
					fmt.Println("Terlalu kecil!")
				} else if guess > secret {
					fmt.Println("Terlalu Besar!")
				} else {
					fmt.Println("Benar! kamu berhasil!!")
					break
				}
}
}