package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

// ThumbnailResult thumbnail işleminin sonucunu tutar
type ThumbnailResult struct {
	URL      string
	Success  bool
	Error    error
	FilePath string
	FileSize int64
	Duration time.Duration
}

// createThumbnail tek bir video URL'si için thumbnail oluşturur
func createThumbnail(videoURL string, index int, wg *sync.WaitGroup, results chan<- ThumbnailResult) {
	defer wg.Done()

	startTime := time.Now()
	result := ThumbnailResult{
		URL: videoURL,
	}

	fmt.Printf("[%d] Thumbnail oluşturuluyor: %s\n", index, videoURL)

	// Thumbnail klasörü oluştur
	thumbnailDir := "thumbnails"
	os.MkdirAll(thumbnailDir, 0755)

	// Thumbnail dosyası oluştur
	timestamp := time.Now().Format("20060102-150405")
	thumbnailFile := filepath.Join(thumbnailDir, fmt.Sprintf("thumbnail_%d_%s.jpg", index, timestamp))
	result.FilePath = thumbnailFile

	// FFmpeg kullanarak doğrudan URL'den thumbnail oluştur
	cmd := exec.Command("ffmpeg",
		"-i", videoURL, // Doğrudan URL kullan
		"-ss", "2", // 2. saniye (başlangıçtan biraz sonra)
		"-vframes", "1", // Tek frame
		"-vf", "scale=320:-1", // Genişlik 320px, yükseklik otomatik
		"-y",                 // Üzerine yaz
		"-loglevel", "quiet", // FFmpeg çıktısını azalt
		thumbnailFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		result.Error = fmt.Errorf("FFmpeg komutu başarısız: %v, çıktı: %s", err, string(output))
		result.Duration = time.Since(startTime)
		results <- result
		return
	}

	// Dosya boyutunu kontrol et
	fileInfo, err := os.Stat(thumbnailFile)
	if err != nil {
		result.Error = fmt.Errorf("Thumbnail dosyası bilgileri alınamadı: %v", err)
		result.Duration = time.Since(startTime)
		results <- result
		return
	}

	result.Success = true
	result.FileSize = fileInfo.Size()
	result.Duration = time.Since(startTime)

	fmt.Printf("[%d] Başarılı! Thumbnail: %s (Boyut: %d bayt, Süre: %v)\n",
		index, thumbnailFile, result.FileSize, result.Duration)

	results <- result
}

func main() {
	// 20 örnek video URL'si (aynı URL'yi kullanıyoruz, farklı URL'lerinizi buraya ekleyin)
	videoURLs := []string{
		"https://n1.coomer.st/data/74/fe/74fee603c3490078efc28fedb4426e57ac52ec3cf482136dc1592710307c8565.mp4",
		"https://n3.coomer.st/data/0c/67/0c677986af3baf74fdca402b8ba18b95038ee9bb88374ba8110235db87e8fda2.m4v",
		"https://n1.coomer.st/data/e9/e7/e9e7fe5e2023d6900a9d9dc30d39c1e936b1196ea579e664230cd9036b2b4cc3.m4v",
		"https://n3.coomer.st/data/19/e3/19e31dc3e36a16826becaec7d80309bb0db3fcd489dc9d89e97c11c9977c135b.mp4",
	}

	fmt.Printf("Toplam %d video için thumbnail oluşturulacak...\n", len(videoURLs))
	startTime := time.Now()

	// WaitGroup ve result channel oluştur
	var wg sync.WaitGroup
	results := make(chan ThumbnailResult, len(videoURLs))

	// Her URL için goroutine başlat
	for i, url := range videoURLs {
		wg.Add(1)
		go createThumbnail(url, i+1, &wg, results)
	}

	// Tüm goroutine'lerin bitmesini bekle
	wg.Wait()
	close(results)

	// Sonuçları topla
	var successCount, errorCount int
	var totalSize int64

	fmt.Println("\n=== ÖZET RAPOR ===")
	for result := range results {
		if result.Success {
			successCount++
			totalSize += result.FileSize
			fmt.Printf("✅ [%s] - %d bayt (%v)\n",
				filepath.Base(result.FilePath), result.FileSize, result.Duration)
		} else {
			errorCount++
			fmt.Printf("❌ [%s] - Hata: %v (%v)\n",
				result.URL, result.Error, result.Duration)
		}
	}

	totalDuration := time.Since(startTime)

	fmt.Printf("\n=== İSTATİSTİKLER ===\n")
	fmt.Printf("Toplam işlem: %d\n", len(videoURLs))
	fmt.Printf("Başarılı: %d\n", successCount)
	fmt.Printf("Başarısız: %d\n", errorCount)
	fmt.Printf("Toplam boyut: %d bayt (%.2f MB)\n", totalSize, float64(totalSize)/1024/1024)
	fmt.Printf("Toplam süre: %v\n", totalDuration)
	fmt.Printf("Ortalama süre: %v\n", totalDuration/time.Duration(len(videoURLs)))

	fmt.Println("\nProgram başarıyla tamamlandı.")
}
