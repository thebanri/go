package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	count := 6000
	var wg sync.WaitGroup

	for i := 1; i < count; i++ {
		port := ":" + strconv.Itoa(i)
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			isOpen(p)
		}(port)
	}

	wg.Wait() // Tüm taramalar bitene kadar bekle
}

func isOpen(port string) {
	timeout := 1 * time.Second
	server, err := net.DialTimeout("tcp", port, timeout)
	if err != nil {
		return
	}

	defer server.Close()
	fmt.Println("Port: ", port)

}
