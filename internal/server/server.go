package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Server struct {
	Handler http.Handler
}

func New() Server {

	mux := http.NewServeMux()

	RegisterRoutes(mux)

	s := Server{
		Handler: mux,
	}

	return s
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /payments", postPayments)
	mux.HandleFunc("GET /payments-summary", getPaymentsSummary)
	mux.HandleFunc("POST /purge-payments", postPurgePayments)
}

type postPaymentsRequest struct {
	Amount        float64 `json:"amount"`
	CorrelationId string  `json:"correlationId"`
}

func postPayments(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		log.Printf("Error reading body: %v", err)
		return
	}

	var parsedRequest postPaymentsRequest

	if err := json.Unmarshal([]byte(body), &parsedRequest); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		log.Printf("Error parsing body: %v", err)
		return
	}

	fmt.Printf("Correlation id: %s, amount: %f\n", parsedRequest.CorrelationId, parsedRequest.Amount)

	if !isCorrelationIdValid(parsedRequest.CorrelationId) {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Write([]byte("You are creating a payment"))
}

func isCorrelationIdValid(id string) bool {

	if err := uuid.Validate(id); err != nil {
		return false
	}

	return true
}

func getPaymentsSummary(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are getting the payments summary"))
}

func postPurgePayments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You just purged all payments in the database"))
}
