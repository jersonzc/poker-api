package main

import (
	"log"
	"net/http"
	"poker-api/poker"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &poker.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":8080", server))
}
