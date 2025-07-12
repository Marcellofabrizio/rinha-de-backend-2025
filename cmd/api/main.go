package main

import (
	"fmt"
	"log"
	"net/http"
	"rinhabackend/internal/server"
)

const port = 42069

func main() {

	s := server.New()

	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s.Handler))
}
