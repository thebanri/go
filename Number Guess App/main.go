package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var guess int

	real_num := rand.IntN(100)
	fmt.Println("Sayı oluşturuldu lütfen tahmin yapınız!")
	for real_num != guess {
		fmt.Printf("Tahmin: ")
		fmt.Scan(&guess)
		if guess > real_num {
			fmt.Println("Sayıyı Azaltın!")
		} else if guess < real_num {
			fmt.Println("Sayıyı Arttırın!")
		} else if guess == real_num {
			fmt.Println("Tebrikler doğru bildiniz! Doğru sayı:", real_num)
			break
		}
	}

}
