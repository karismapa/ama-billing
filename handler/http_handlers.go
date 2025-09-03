package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/karismapa/ama-billing/model"
)

func (s *LoanHTTPServer) createLoan(w http.ResponseWriter, r *http.Request) {
	// read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error: ReadAll failed, err: %v", err)
		http.Error(w, "error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON into struct
	var payload model.LoanDisplay
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("error: Unmarshal failed, err: %v", err)
		http.Error(w, "error parsing JSON: "+err.Error(), http.StatusBadRequest)
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
	vars := mux.Vars(r)
	loanIDStr := vars["id"]

	loanID, err := strconv.ParseInt(loanIDStr, 10, 64)
	if err != nil {
		log.Printf("error: ParseInt failed, err: %v", err)
		http.Error(w, "invalid loan ID", http.StatusBadRequest)
		return
	}

	loan, err := s.loanUsecase.GetLoan(r.Context(), loanID)
	if err != nil {
		log.Printf("error: GetLoan failed, err: %v", err)
		http.Error(w, "error on getting loan info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.PackLoanDisplay(loan))
}

func (s *LoanHTTPServer) getOutstandingInstallments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanIDStr := vars["id"]

	loanID, err := strconv.ParseInt(loanIDStr, 10, 64)
	if err != nil {
		log.Printf("error: ParseInt failed, err: %v", err)
		http.Error(w, "invalid loan ID", http.StatusBadRequest)
		return
	}

	outstandings, err := s.loanUsecase.GetOutstandingInstallments(r.Context(), loanID)
	if err != nil {
		log.Printf("error: GetLoan failed, err: %v", err)
		http.Error(w, "error on getting loan info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.PackInstallmentsDisplay(outstandings))
}

func (s *LoanHTTPServer) payInstallment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanIDStr := vars["id"]

	loanID, err := strconv.ParseInt(loanIDStr, 10, 64)
	if err != nil {
		log.Printf("error: ParseInt failed, err: %v", err)
		http.Error(w, "invalid loan ID", http.StatusBadRequest)
		return
	}

	err = s.loanUsecase.PayInstallment(r.Context(), loanID)
	if err != nil {
		log.Printf("error: PayInstallment failed, err: %v", err)
		http.Error(w, "error on pay process", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
