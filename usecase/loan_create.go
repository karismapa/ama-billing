package usecase

import (
	"context"
	"time"

	"github.com/karismapa/ama-billing/model"
	"github.com/karismapa/ama-billing/utils"
)

func (u *LoanUsecase) CreateLoan(ctx context.Context, loan model.Loan) (generatedLoan *model.Loan, err error) {
	// input checking
	if loan.UserID < utils.MinUserID {
		return nil, utils.ErrInvalidUserID
	}
	if loan.PrincipalValue < utils.MinPrincipalValue {
		return nil, utils.ErrPrincipalValueTooSmall
	}
	if loan.PrincipalValue > utils.MaxPrincipalValue {
		return nil, utils.ErrPrincipalValueTooBig
	}
	if loan.Rate < utils.MinRateValue {
		return nil, utils.ErrRateValueTooSmall
	}
	if loan.Rate > utils.MaxRateValue {
		return nil, utils.ErrRateValueTooBig
	}
	if loan.NumOfInstallment < utils.MinNumOfInstallment {
		return nil, utils.ErrNumOfInstallmentTooSmall
	}
	if loan.NumOfInstallment > utils.MaxNumOfInstallment {
		return nil, utils.ErrNumOfInstallmentTooBig
	}

	// assign predefined values
	loan.Status = model.LoanStatusActive // skip initial, directly assign active status
	calculateInstallment(&loan)

	// create and return
	return u.loanInmem.CreateLoan(ctx, loan)
}

func calculateInstallment(loan *model.Loan) {
	loan.Installments = []*model.LoanInstallment{}
	principalValue, interestValue := calculateInstallmentValues(loan.PrincipalValue, loan.Rate, loan.NumOfInstallment)
	dueTimeUnixes := getInstallmentDueTimeUnixes(time.Now(), loan.NumOfInstallment)
	for _, dueTimeUnix := range dueTimeUnixes {
		loan.Installments = append(loan.Installments, &model.LoanInstallment{
			PrincipalValue: principalValue,
			InterestValue:  interestValue,
			DueTimeUnix:    dueTimeUnix,
			Status:         model.LoanInstallmentStatusPending, // skip initial, directly assign pending status
		})
	}
}

func calculateInstallmentValues(loanPrincipalValue int64, loanAnnualRate int32, loanNumOfInstallment int32) (installmentPrincipalValue int64, loanInterestValue int64) {
	installmentPrincipalValue = loanPrincipalValue / int64(loanNumOfInstallment)
	loanInterestValue = installmentPrincipalValue * int64(loanAnnualRate) / 10000
	return
}

func getInstallmentDueTimeUnixes(startTime time.Time, loanNumOfInstallment int32) (dueTimeUnixes []int64) {
	dueTimeUnixes = make([]int64, 0)
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
	for i := range loanNumOfInstallment {
		dueTimeUnix := startTime.Add(time.Duration(i+1) * 7 * 24 * time.Hour).Unix()
		dueTimeUnixes = append(dueTimeUnixes, dueTimeUnix)
	}
	return
}
