package usecase

import (
	"context"

	"github.com/karismapa/ama-billing/model"
	"github.com/karismapa/ama-billing/repository/inmem"
)

type ILoanUsecase interface {
	CreateLoan(ctx context.Context, loan model.Loan) (generatedLoan *model.Loan, err error)
	GetOutstandingInstallments(ctx context.Context, loanID int64) (installments []*model.LoanInstallment, err error)
}

type LoanUsecase struct {
	loanInmem inmem.ILoanInmem
}

func NewLoanUsecase() ILoanUsecase {
	return &LoanUsecase{
		loanInmem: inmem.NewLoanInmem(),
	}
}
