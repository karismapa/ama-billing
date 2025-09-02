package inmem

import (
	"context"
	"sync"
	"time"

	"github.com/karismapa/ama-billing/model"
	"github.com/karismapa/ama-billing/utils"
)

var ()

type ILoanInmem interface {
	CreateLoan(ctx context.Context, loan model.Loan) (generatedLoan *model.Loan, err error)
	GetLoan(ctx context.Context, loanID int64) (loan model.Loan, err error)
	GetInstallments(ctx context.Context, loanID int64, status *model.LoanInstallmentStatus, endDueTimeUnix *int64) (installments []*model.LoanInstallment, err error)
	GetOldestInstallment(ctx context.Context, loanID int64, status *model.LoanInstallmentStatus) (oldestInstallment *model.LoanInstallment, err error)
	UpdateInstallmentStatus(ctx context.Context, installmentID int64, status model.LoanInstallmentStatus) (err error)
}

type LoanInmem struct {
	loanTable            []*model.Loan
	loanMap              map[int64]*model.Loan
	loanInstallmentTable []*model.LoanInstallment
	loanInstallmentMap   map[int64]*model.LoanInstallment

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
	m.loanMap[generatedLoan.ID] = generatedLoan
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
		m.loanInstallmentMap[generatedLoanInstallment.ID] = generatedLoanInstallment
	}
	m.mu.Unlock()

	generatedLoan.Installments = generatedLoanInstallments
	return
}

func (m *LoanInmem) GetInstallments(ctx context.Context, loanID int64, status *model.LoanInstallmentStatus, endDueTimeUnix *int64) (installments []*model.LoanInstallment, err error) {
	for _, installment := range m.loanInstallmentTable {
		if installment == nil {
			continue
		}
		if installment.LoanID != loanID {
			continue
		}
		if status != nil && installment.Status != *status {
			continue
		}
		if endDueTimeUnix != nil && installment.DueTimeUnix > *endDueTimeUnix {
			continue
		}
		installments = append(installments, installment)
	}
	return
}

func (m *LoanInmem) GetOldestInstallment(ctx context.Context, loanID int64, status *model.LoanInstallmentStatus) (oldestInstallment *model.LoanInstallment, err error) {
	installments, err := m.GetInstallments(ctx, loanID, status, nil)
	if err != nil {
		return nil, err
	}
	if len(installments) == 0 {
		return nil, utils.ErrInstallmentNotFound
	}
	for _, selection := range installments {
		if selection == nil {
			continue
		}
		if oldestInstallment == nil || oldestInstallment.DueTimeUnix > selection.DueTimeUnix {
			oldestInstallment = selection
		}
	}
	return
}

func (m *LoanInmem) UpdateInstallmentStatus(ctx context.Context, installmentID int64, status model.LoanInstallmentStatus) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	installment, ok := m.loanInstallmentMap[installmentID]
	if !ok {
		return utils.ErrInstallmentNotFound
	}
	installment.Status = status
	return
}
