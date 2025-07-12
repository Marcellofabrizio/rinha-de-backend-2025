package server

import "net/http"

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

func postPayments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are creating a payment"))
}

func getPaymentsSummary(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are getting the payments summary"))
}

func postPurgePayments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You just purged all payments in the database"))
}
