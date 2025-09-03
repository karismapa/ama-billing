package model

type LoanDisplay struct {
	ID                int64      `json:"id"`
	UserID            int64      `json:"user_id"`
	PrincipalValue    int64      `json:"principal_value"`
	PrincipalValueStr string     `json:"principal_value_str"`
	Rate              int32      `json:"rate"`
	RateStr           string     `json:"rate_str"`
	CreateTimeUnix    int64      `json:"create_time_unix"`
	UpdateTimeUnix    int64      `json:"update_time_unix"`
	Status            LoanStatus `json:"loan_status"`
	NumOfInstallment  int32      `json:"num_of_installment"`

	Installments []LoanInstallmentDisplay `json:"installments"`
}

type LoanInstallmentDisplay struct {
	ID                int64                 `json:"id"`
	LoanID            int64                 `json:"loan_id"`
	PrincipalValue    int64                 `json:"principal_value"` // in rupiah *100
	PrincipalValueStr string                `json:"principal_value_str"`
	InterestValue     int64                 `json:"interest_value"` // in rupiah *100
	InterestValueStr  string                `json:"interest_value_str"`
	DueTimeUnix       int64                 `json:"due_time_unix"`
	CreateTimeUnix    int64                 `json:"create_time_unix"`
	UpdateTimeUnix    int64                 `json:"update_time_unix"`
	Status            LoanInstallmentStatus `json:"status"`
}
