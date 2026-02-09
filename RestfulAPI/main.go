package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Just GET Method have access")) //bookmark

		return
	}

	w.Write([]byte("GET isteği başarıyla alındı!"))
}

func main() {

	http.HandleFunc("/monster", Handler)

	err := http.ListenAndServe(":9090", nil)
	fmt.Println(err)
}
