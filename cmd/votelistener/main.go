package main

import (
	crypto "crypto/rand"
	"encoding/binary"
	"github.com/TicketsBot/VoteListener/config"
	"github.com/TicketsBot/VoteListener/database"
	"github.com/TicketsBot/VoteListener/http"
	"io"
	"math/rand"
)

func main() {
	// seed random
	b := make([]byte, 8)
	if _, err := io.ReadFull(crypto.Reader, b); err != nil {
		panic(err)
	}

	rand.Seed(int64(binary.LittleEndian.Uint64(b)))

	config.LoadConfig()
	database.ConnectDatabase()
	http.StartServer()
}
