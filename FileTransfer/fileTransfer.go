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

	for {

		conn, err := net.Dial("tcp", ":3050")
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}

		reader := bufio.NewReader(os.Stdin)

		fileName, _ := reader.ReadString('\n')
		file, err := os.Open(strings.TrimSpace(fileName))
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}

		fmt.Fprintln(conn, fileName)

		_, err = io.Copy(conn, file)
		file.Close()
		conn.Close()
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}

		fmt.Println("Dosya gönderildi!")

	}

}
