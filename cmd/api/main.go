package main

import (
	"fmt"
	"log"
	"net/http"
	"rinhabackend/internal/db"
	"rinhabackend/internal/server"
)

const port = 42069

func main() {

	s := server.New()

	db.Init("localhost:6379")

	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s.Handler))
}
