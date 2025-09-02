package usecase

import (
	"context"
	"time"

	"github.com/karismapa/ama-billing/model"
	"github.com/karismapa/ama-billing/utils"
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

func (u *LoanUsecase) PayInstallment(ctx context.Context, loanID int64) (err error) {
	filterStatus := model.LoanInstallmentStatusPending
	oldestInstallment, err := u.loanInmem.GetOldestInstallment(ctx, loanID, &filterStatus)
	if err != nil {
		return err
	}
	if oldestInstallment.Status != model.LoanInstallmentStatusPending {
		return utils.ErrInstallmentStatusInvalid
	}
	return u.loanInmem.UpdateInstallmentStatus(ctx, oldestInstallment.ID, model.LoanInstallmentStatusPaid)
}
