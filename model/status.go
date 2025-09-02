package model

type LoanStatus int32

const (
	LoanStatusInitial LoanStatus = iota // for drafting, move to active after approved
	LoanStatusActive
	LoanStatusOverdue
	LoanStatusPaidOff
	LoanStatusClosed
)

type LoanInstallmentStatus int32

const (
	LoanInstallmentStatusInitial LoanInstallmentStatus = iota // for drafting, move to pending after approved
	LoanInstallmentStatusPending
	LoanInstallmentStatusPaid
)
