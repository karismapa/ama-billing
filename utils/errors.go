package utils

type ErrorConst string

func (e ErrorConst) Error() string {
	return string(e)
}

const (
	ErrInvalidUserID            ErrorConst = "invalid user ID"
	ErrPrincipalValueTooSmall   ErrorConst = "principal value too small"
	ErrPrincipalValueTooBig     ErrorConst = "principal value too big"
	ErrRateValueTooSmall        ErrorConst = "rate value too small"
	ErrRateValueTooBig          ErrorConst = "rate value too big"
	ErrNumOfInstallmentTooSmall ErrorConst = "num of installment too small"
	ErrNumOfInstallmentTooBig   ErrorConst = "num of installment too big"

	ErrInstallmentStatusInvalid ErrorConst = "installment status invalid"
)

const (
	ErrLoanNotFound        ErrorConst = "loan not found"
	ErrInstallmentNotFound ErrorConst = "installment not found"
)
