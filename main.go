package main

import (
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	//file, header, err := r.FormFile("file")

}

func main() {

	http.HandleFunc("/upload", uploadHandler)
	_ = http.ListenAndServe(":8080", nil)

}
