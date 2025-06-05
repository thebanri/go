package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type CatData struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func DownloadFile(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	out, err := os.Create("saved.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)

}

func GenerateUrl() []byte {
	baseURL := "https://api.thecatapi.com/v1/images/search"
	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in body:", err)
	}
	return body
}

func main() {

	for {
		body := GenerateUrl()
		var cats []CatData
		json.Unmarshal(body, &cats)

		var url string
		if len(cats) > 0 {
			url = cats[0].Url
		} else {
			fmt.Println("Url verisi dönmedi!")
			continue
		}

		DownloadFile(url)

		time.Sleep(2 * time.Second)

	}

}
