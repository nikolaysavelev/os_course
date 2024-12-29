package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const fileStorage = "/fileStorage"

func main() {
	http.HandleFunc("/create", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Request from client received")
		fileName := req.URL.Query().Get("filename")
		if fileName == "" {
			http.Error(w, "Missing query parameter", http.StatusBadRequest)
			return
		}

		filePath := filepath.Join(fileStorage, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create file: %v", err), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		w.WriteHeader(http.StatusCreated)
		fmt.Printf("File '%s' created successfully\n", fileName)
	})

	if err := os.MkdirAll(fileStorage, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}