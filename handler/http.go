package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karismapa/ama-billing/usecase"
)

type ILoanHTTPServer interface {
	Serve()
}

type LoanHTTPServer struct {
	loanUsecase usecase.ILoanUsecase
	router      *mux.Router
}

func NewHTTPServer() ILoanHTTPServer {
	server := &LoanHTTPServer{
		loanUsecase: usecase.NewLoanUsecase(),
	}
	server.registerRoutes()
	return server
}

func (s *LoanHTTPServer) registerRoutes() {
	s.router = mux.NewRouter()

	s.router.HandleFunc("/loan", s.createLoan).Methods("POST")
	s.router.HandleFunc("/loan/{id}", s.getLoanDetail).Methods("GET")
	s.router.HandleFunc("/loan/{id}/outstandings", s.getOutstandingInstallments).Methods("GET") // GetOutstanding
	s.router.HandleFunc("/loan/{id}/outstanding_recap", s.getOutstandingRecap).Methods("GET")   // GetTotalOutstanding and IsDelinquent
	s.router.HandleFunc("/loan/{id}/pay", s.payInstallment).Methods("POST")                     // MakePayment
}

func (s *LoanHTTPServer) Serve() {
	// Start the server on port 8080
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
