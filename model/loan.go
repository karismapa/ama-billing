package model

type Loan struct {
	ID               int64 // loan ID
	UserID           int64
	PrincipalValue   int64 // in rupiah *100
	Rate             int32 // in percentage *100
	CreateTimeUnix   int64
	UpdateTimeUnix   int64
	Status           LoanStatus
	NumOfInstallment int32

	Installments []*LoanInstallment
}

type LoanStatus int32

const (
	LoanStatusInitial LoanStatus = iota // for drafting, move to active after approved
	LoanStatusActive
	LoanStatusOverdue
	LoanStatusPaidOff
	LoanStatusClosed
)

type LoanInstallment struct {
	ID             int64
	LoanID         int64
	PrincipalValue int64 // in rupiah *100
	InterestValue  int64 // in rupiah *100
	DueTimeUnix    int64
	CreateTimeUnix int64
	UpdateTimeUnix int64
	Status         LoanInstallmentStatus
}

type LoanInstallmentStatus int32

const (
	LoanInstallmentStatusInitial LoanInstallmentStatus = iota // for drafting, move to pending after approved
	LoanInstallmentStatusPending
	LoanInstallmentStatusPaid
)
