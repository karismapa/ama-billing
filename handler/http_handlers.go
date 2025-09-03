package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/karismapa/ama-billing/model"
)

func (s *LoanHTTPServer) createLoan(w http.ResponseWriter, r *http.Request) {
	// read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON into struct
	var payload model.LoanDisplay
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("error: Unmarshal failed, err: %v", err)
		http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("create loan with payload: %v", payload)

	result, err := s.loanUsecase.CreateLoan(r.Context(), model.Loan{
		UserID:           payload.UserID,
		PrincipalValue:   payload.PrincipalValue,
		Rate:             payload.Rate,
		NumOfInstallment: payload.NumOfInstallment,
	})
	if err != nil {
		log.Printf("error: CreateLoan failed, err: %v", err)
		http.Error(w, "error on creating new loan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.PackLoanDisplay(result))
}

func (s *LoanHTTPServer) getLoanDetail(w http.ResponseWriter, r *http.Request) {

}

func (s *LoanHTTPServer) getOutstandingInstallments(w http.ResponseWriter, r *http.Request) {

}

func (s *LoanHTTPServer) payInstallment(w http.ResponseWriter, r *http.Request) {

}
