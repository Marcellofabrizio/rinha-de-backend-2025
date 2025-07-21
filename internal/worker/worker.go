package worker

import (
	"encoding/json"
	"fmt"
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
	log.Printf("Starting worker %d", w.WorkerId)
	for {
		paymentString, err := db.Client.RPopLPush(db.DbCtx, db.PendingQueue, w.ProcessingQueue).Result()

		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("[%d] - pending content: %s\n", w.WorkerId, paymentString)

		var payment server.Payment

		if err := json.Unmarshal([]byte(paymentString), &payment); err != nil {
			log.Printf("[%d] - error parsing body: %v\n", w.WorkerId, err)
			time.Sleep(5 * time.Second)
			continue
		}

	}

}
