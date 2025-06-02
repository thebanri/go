package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	server, err := net.Listen("tcp", ":9000") // 9000 portunu dinlemeye başlar
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	defer server.Close()

	for {

		conn, err := server.Accept() //O porta bağlanmasını sağlar bağlanırsa devam eder
		if err != nil {
			fmt.Println("Hata:", err)
			continue
		}

		go handleClient(conn)

	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Bağlantı kapandı veya hata:", err)
			break
		}

		message = strings.TrimSpace(message)
		if message == "quit" {
			fmt.Println("İstemci çıkış yaptı.")
			break
		}

		fmt.Println("Gelen Mesaj:", message)
	}
}
