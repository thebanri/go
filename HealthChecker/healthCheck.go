package main

import (
	"fmt"
	"net/http"
	"time"
)

func healthCheck(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Hata: ", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println(url, "\nSunucu Çalışıyor!")
	} else {
		fmt.Println(url, "Hata: ")
	}

}

func main() {

	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		go healthCheck("http://localhost:5000")
	}

}
