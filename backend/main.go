package main

import (
	"log"

	"github.com/actatum/url-shortener/backend/shortener"
)

func main() {
	if err := shortener.Run(); err != nil {
		log.Fatal(err)
	}
}
