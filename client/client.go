package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	url := "http://server:8080/create?filename=osfile.txt"
	response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("RESPONSE STATUS = %s\n", response.Status)
}