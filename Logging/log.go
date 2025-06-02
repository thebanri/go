package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	server, err := http.Get("http://localhost:5000")
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}

	defer server.Body.Close()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	slog.Info("Sunucuya bağlandı!", "response:", server.Status)
}
