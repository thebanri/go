package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	server, err := net.Listen("tcp", ":3050")
	if err != nil {

		fmt.Println("Hata:", err)
		return
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {

			fmt.Println("Hata:", err)
			continue
		}

		go takeFile(conn)
	}

}

func takeFile(conn net.Conn) {

	defer conn.Close()
	reader := bufio.NewReader(conn)

	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Dosya İsim hatası:", err)
		return
	}

	fileName = strings.TrimSpace(fileName)
	outputFilename := "copy_" + fileName

	file, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Dosya Oluşturma Hatası:", err)
		return
	}

	defer file.Close()

	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Dosya İçeriği Alma Hatası:", err)
		return
	} else {
		fmt.Println("Dosya içeriği başarıyla alındı!")
	}

}
