package main

import (
	"fmt"
	"log"
	"net/http"
	"rinhabackend/internal/db"
	"rinhabackend/internal/server"
	"rinhabackend/internal/worker"
)

func main() {

	s := server.New()

	dbHost := "127.0.0.1"// os.Getenv("REDIS_HOST")
	dbPort := "6379" // os.Getenv("REDIS_PORT")

	db.Init(fmt.Sprintf("%s:%s", dbHost, dbPort))

	port := "42069" // os.Getenv("PORT")

	concurrentWorkers := 2

	for wId := 0; wId < concurrentWorkers; wId++ {

		go func() {
			log.Println("Booting up worker")
			w := worker.Worker{
				WorkerId:        wId,
				ProcessingQueue: db.ProcessingQueue,
			}

			w.StartWorker()
		}()

	}

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s.Handler))
}
