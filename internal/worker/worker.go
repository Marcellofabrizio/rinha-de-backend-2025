package worker

import (
	"encoding/json"
	"log"
	"rinhabackend/internal/db"
	"rinhabackend/internal/server"
	"time"
)

type Status int

const (
	Pending Status = iota
	Failed
	Processed
	Fallback
)

type PaymentTask struct {
	Content string
	Gateway string
	Status  Status
	Value   float64
}

type Worker struct {
	WorkerId        int
	ProcessingQueue string
}

func (w *Worker) StartWorker() {

	for {
		paymentString, err := db.Client.RPopLPush(db.DbCtx, db.PendingQueue, w.ProcessingQueue).Result()

		if err != nil {
			log.Fatal(err)
			time.Sleep(5 * time.Second)
			continue
		}

		var payment server.Payment

		if err := json.Unmarshal([]byte(paymentString), &payment); err != nil {
			log.Printf("Error parsing body: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
	}

}
