package inmem

import (
	"context"
	"sync"
	"time"

	"github.com/karismapa/ama-billing/model"
)

var ()

type ILoanInmem interface {
	CreateLoan(ctx context.Context, loan model.Loan) (generatedLoan *model.Loan, err error)
	GetLoan(ctx context.Context, loanID int64) (loan model.Loan, err error)
}

type LoanInmem struct {
	loanTable            []*model.Loan
	loanInstallmentTable []*model.LoanInstallment

	mu sync.Mutex
}

func (m *LoanInmem) CreateLoan(ctx context.Context, loan model.Loan) (generatedLoan *model.Loan, err error) {
	generatedLoan = &model.Loan{
		UserID:           loan.UserID,
		PrincipalValue:   loan.PrincipalValue,
		Rate:             loan.Rate,
		Status:           loan.Status,
		NumOfInstallment: loan.NumOfInstallment,
	}

	// store to inmem
	m.mu.Lock()
	loanID := int64(len(m.loanTable)) + 1
	actionTime := time.Now()
	generatedLoan.ID = loanID
	generatedLoan.CreateTimeUnix = actionTime.Unix()
	generatedLoan.UpdateTimeUnix = actionTime.Unix()
	m.loanTable = append(m.loanTable, generatedLoan)
	m.mu.Unlock()

	generatedLoanInstallments := []*model.LoanInstallment{}
	for _, loanInstallment := range loan.Installments {
		generatedLoanInstallments = append(generatedLoanInstallments, &model.LoanInstallment{
			LoanID:         generatedLoan.ID,
			PrincipalValue: loanInstallment.PrincipalValue,
			InterestValue:  loanInstallment.InterestValue,
			DueTimeUnix:    loanInstallment.DueTimeUnix,
			Status:         loanInstallment.Status,
		})
	}

	m.mu.Lock()
	loanInstallmentID := int64(len(m.loanInstallmentTable)) + 1
	actionTime = time.Now()
	for i, generatedLoanInstallment := range generatedLoanInstallments {
		generatedLoanInstallment.ID = loanInstallmentID + int64(i)
		generatedLoanInstallment.CreateTimeUnix = actionTime.Unix()
		generatedLoanInstallment.UpdateTimeUnix = actionTime.Unix()
		m.loanInstallmentTable = append(m.loanInstallmentTable, generatedLoanInstallment)
	}
	m.mu.Unlock()

	generatedLoan.Installments = generatedLoanInstallments
	return
}
