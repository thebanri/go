package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	client, err := net.Dial("tcp", ":9000")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	defer client.Close()

	writer := bufio.NewWriter(client)
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Mesaj gönder:")
		text, _ := reader.ReadString('\n') // burada \n alınır
		text = strings.TrimSpace(text) + "\n"
		writer.WriteString(text)
		writer.Flush()

		fmt.Println("Mesaj gönderildi")
	}

}
