package main

import (
	"log"
	"net/http"
	"poker-api/poker"
)

func main() {
	server := &poker.PlayerServer{}
	log.Fatal(http.ListenAndServe(":8080", server))
}
