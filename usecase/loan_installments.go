package usecase

import (
	"context"
	"log"
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
	filterTime := calculateInstallmentFilterTime()
	return u.loanInmem.GetInstallments(ctx, loanID, &filterStatus, &filterTime)
}

func (u *LoanUsecase) GetOutstandingRecap(ctx context.Context, loanID int64) (outstandingRecap model.OutstandingRecap, err error) {
	filterTime := calculateInstallmentFilterTime()
	totalOutstandingValue, installmentCount, err := u.loanInmem.GetTotalOutstanding(ctx, loanID, &filterTime)
	if err != nil {
		log.Printf("GetTotalOutstanding failed, err: %v", err)
		return
	}
	outstandingRecap.TotalOutstandingValue = totalOutstandingValue
	outstandingRecap.IsDelinquent = installmentCount > utils.MaxInstallmentForDelinquent
	return
}

func calculateInstallmentFilterTime() (endTimeUnix int64) {
	return time.Now().Add(7 * 24 * time.Hour).Unix()
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
