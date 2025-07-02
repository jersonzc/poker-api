package main

import (
	"log"
	"net/http"
	"poker-api/poker"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8080", server))
}
