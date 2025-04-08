package main

import (
	"fmt"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(header.Filename)
}

func main() {

	http.HandleFunc("/upload", uploadHandler)
	_ = http.ListenAndServe(":8080", nil)

}
