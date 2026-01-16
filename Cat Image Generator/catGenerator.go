package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

// CatData API'den dönen JSON yapısını temsil eder
type CatData struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

// DownloadFile resmi atomik bir şekilde indirir ve günceller
func DownloadFile(url string) {
	// 1. Yeni dosyanın uzantısını ve adını belirle
	ext := path.Ext(url)
	if ext == "" {
		ext = ".jpg"
	}
	finalName := "current_cat" + ext
	tempName := "downloading_temp"

	// 2. İndirme işlemini başlat
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer resp.Body.Close()

	// 3. Önce geçici dosyaya yaz (Zed'in çökmesini önlemek için şart)
	out, err := os.Create(tempName)
	if err != nil {
		return
	}
	_, err = io.Copy(out, resp.Body)
	out.Close()
	if err != nil {
		return
	}

	// 4. TEMİZLİK: Klasördeki eski 'current_cat' dosyalarını sil
	// Böylece klasörde sadece tek bir dosya kalacak
	oldFiles := []string{"current_cat.jpg", "current_cat.png", "current_cat.gif", "current_cat.jpeg"}
	for _, f := range oldFiles {
		if f != finalName { // Eğer yeni dosya adıyla eski dosya adı aynı değilse sil
			os.Remove(f)
		}
	}

	// 5. ATOMİK RENAME: Geçici dosyayı asıl ismine çevir
	err = os.Rename(tempName, finalName)
	if err != nil {
		fmt.Println("Rename hatası:", err)
	} else {
		fmt.Printf("[%s] Güncellendi: %s\n", time.Now().Format("15:04:05"), finalName)
	}
}

// GenerateUrl API'den rastgele kedi verisini çeker
func GenerateUrl() []byte {
	baseURL := "https://api.thecatapi.com/v1/images/search"
	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("API bağlantı hatası:", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return body
}

func main() {
	fmt.Println("CatGenerator başlatıldı... (Çıkmak için Ctrl+C)")

	for {
		body := GenerateUrl()
		if body == nil {
			time.Sleep(5 * time.Second)
			continue
		}

		var cats []CatData
		err := json.Unmarshal(body, &cats)

		if err == nil && len(cats) > 0 {
			DownloadFile(cats[0].Url)
		} else {
			fmt.Println("Veri işlenemedi veya boş döndü.")
		}

		// Zed'in dosya sistemini çok sık tarayıp yorulmaması için
		// 5 saniyelik bir bekleme süresi idealdir.
		time.Sleep(1 * time.Second)
	}
}
