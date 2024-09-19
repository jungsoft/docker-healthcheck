package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url_env := os.Getenv("URL")
	port_env := os.Getenv("PORT")
	var url string

	if url_env != "" {
		url = url_env
	} else {
		if port_env != "" {
			url = "http://localhost:" + port_env
		} else {
			url = "http://localhost"
		}
	}

	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
