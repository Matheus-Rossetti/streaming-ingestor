package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
		return
	}

	if header.Header.Get("Content-Type") != "video/mp4" {
	}

	fmt.Println(header.Filename)
	fmt.Println(header.Size)
	fmt.Println(header.Header.Get("Content-Type"))

}

func main() {

	http.HandleFunc("/upload", uploadHandler)
	_ = http.ListenAndServe(":8080", nil)

}
