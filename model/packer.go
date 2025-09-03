package model

import (
	"fmt"
	"strings"
)

func PackLoanDisplay(loan *Loan) (loanDisplay LoanDisplay) {
	if loan == nil {
		return
	}
	loanDisplay = LoanDisplay{
		ID:                loan.ID,
		UserID:            loan.UserID,
		PrincipalValue:    loan.PrincipalValue,
		PrincipalValueStr: formatRupiahWithDecimals(loan.PrincipalValue),
		Rate:              loan.Rate,
		RateStr:           formatRateWithDecimals(loan.Rate),
		CreateTimeUnix:    loan.CreateTimeUnix,
		UpdateTimeUnix:    loan.UpdateTimeUnix,
		Status:            loan.Status,
		NumOfInstallment:  loan.NumOfInstallment,
	}
	if len(loan.Installments) > 0 {
		loanDisplay.Installments = PackInstallmentsDisplay(loan.Installments)
	}
	return
}

func PackInstallmentsDisplay(installments []*LoanInstallment) (installmentsDisplay []LoanInstallmentDisplay) {
	installmentsDisplay = make([]LoanInstallmentDisplay, 0)
	for _, installment := range installments {
		if installment == nil {
			continue
		}
		installmentsDisplay = append(installmentsDisplay, LoanInstallmentDisplay{
			ID:                installment.ID,
			LoanID:            installment.LoanID,
			PrincipalValue:    installment.PrincipalValue,
			PrincipalValueStr: formatRupiahWithDecimals(installment.PrincipalValue),
			InterestValue:     installment.InterestValue,
			InterestValueStr:  formatRupiahWithDecimals(installment.InterestValue),
			DueTimeUnix:       installment.DueTimeUnix,
			CreateTimeUnix:    installment.CreateTimeUnix,
			UpdateTimeUnix:    installment.UpdateTimeUnix,
			Status:            installment.Status,
		})
	}
	return
}

func PackOutstandingRecap(in OutstandingRecap) OutstandingRecap {
	in.TotalOutstandingValueStr = formatRupiahWithDecimals(in.TotalOutstandingValue)
	return in
}

func formatRupiahWithDecimals(amountInCents int64) string {
	// split into rupiah and cents
	rupiah := amountInCents / 100
	cents := amountInCents % 100

	// format rupiah part
	rupiahStr := fmt.Sprintf("%d", rupiah)
	formattedRupiah := addThousandSeparators(rupiahStr)

	// format with decimals
	result := fmt.Sprintf("Rp %s,%02d", formattedRupiah, cents)

	return result
}

// helper function to add thousand separators (dots)
func addThousandSeparators(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var result strings.Builder

	// add characters from right to left with separators
	for i, char := range s {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteString(".")
		}
		result.WriteRune(char)
	}

	return result.String()
}

func formatRateWithDecimals(rateInCents int32) string {
	// split into rate and cents
	rate := rateInCents / 100
	cents := rateInCents % 100

	// format rate part
	rupiahStr := fmt.Sprintf("%d", rate)

	// format with decimals
	result := fmt.Sprintf("%s,%02d%%", rupiahStr, cents)

	return result
}
