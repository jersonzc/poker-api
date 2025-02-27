package main

import (
	"log"
	"net/http"
	"poker-api/poker"
)

func main() {
	handler := http.HandlerFunc(poker.PlayerServer)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
