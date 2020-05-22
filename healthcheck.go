package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	var url string

	if port != "" {
		url = "http://localhost:" + port
	} else {
		url = "http://localhost"
	}

	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
