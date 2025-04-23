package main

import (
	"log"
	"net/http"
	"poker-api/poker"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := &poker.PlayerServer{store}
	log.Fatal(http.ListenAndServe(":8080", server))
}
