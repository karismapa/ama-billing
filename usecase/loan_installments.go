package usecase

import (
	"context"
	"time"

	"github.com/karismapa/ama-billing/model"
)

func (u *LoanUsecase) GetOutstandingInstallments(ctx context.Context, loanID int64) (installments []*model.LoanInstallment, err error) {
	// outstanding = pending + overdue installments;
	// since we only have status:
	// - INITIAL
	// - PENDING
	// - PAID
	// we can just filter by status=PENDING
	filterStatus := model.LoanInstallmentStatusPending
	filterTime := time.Now().Add(7 * 24 * time.Hour).Unix()
	return u.loanInmem.GetInstallments(ctx, loanID, &filterStatus, &filterTime)
}
