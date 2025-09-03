package usecase

import (
	"context"

	"github.com/karismapa/ama-billing/model"
)

func (u *LoanUsecase) GetLoan(ctx context.Context, loanID int64) (loan *model.Loan, err error) {
	return u.loanInmem.GetLoan(ctx, loanID)
}
