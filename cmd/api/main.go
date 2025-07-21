package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rinhabackend/internal/db"
	"rinhabackend/internal/server"
)

func main() {

	s := server.New()

	dbHost := os.Getenv("REDIS_HOST")
	dbPort := os.Getenv("REDIS_PORT")

	db.Init(fmt.Sprintf("%s:%s", dbHost, dbPort))

	port := os.Getenv("PORT")

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s.Handler))
}
