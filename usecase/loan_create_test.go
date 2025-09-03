package usecase

import (
	"reflect"
	"testing"
	"time"
)

func TestGetInstallmentDueTimeUnixes(t *testing.T) {
	tests := []struct {
		name                  string
		startTime             time.Time
		loanNumOfInstallment  int32
		expectedDueTimeUnixes []int64
	}{
		{
			name:                 "Single installment",
			startTime:            time.Date(2024, 1, 1, 10, 30, 45, 0, time.UTC),
			loanNumOfInstallment: 1,
			expectedDueTimeUnixes: []int64{
				time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC).Unix(), // 1 week later
			},
		},
		{
			name:                 "Multiple installments (3)",
			startTime:            time.Date(2024, 1, 1, 15, 45, 30, 0, time.UTC),
			loanNumOfInstallment: 3,
			expectedDueTimeUnixes: []int64{
				time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC).Unix(),  // 1 week
				time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Unix(), // 2 weeks
				time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC).Unix(), // 3 weeks
			},
		},
		{
			name:                  "Zero installments",
			startTime:             time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			loanNumOfInstallment:  0,
			expectedDueTimeUnixes: []int64{},
		},
		{
			name:                 "Start time with different timezone (should normalize to UTC)",
			startTime:            time.Date(2024, 6, 15, 14, 30, 20, 500000000, time.FixedZone("WIB", 7*3600)), // UTC+7
			loanNumOfInstallment: 2,
			expectedDueTimeUnixes: []int64{
				time.Date(2024, 6, 22, 0, 0, 0, 0, time.UTC).Unix(), // 1 week from June 15
				time.Date(2024, 6, 29, 0, 0, 0, 0, time.UTC).Unix(), // 2 weeks from June 15
			},
		},
		{
			name:                 "Month boundary crossing",
			startTime:            time.Date(2024, 1, 29, 12, 0, 0, 0, time.UTC), // Near end of January
			loanNumOfInstallment: 2,
			expectedDueTimeUnixes: []int64{
				time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC).Unix(),  // February 5
				time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC).Unix(), // February 12
			},
		},
		{
			name:                 "Year boundary crossing",
			startTime:            time.Date(2023, 12, 25, 23, 59, 59, 0, time.UTC),
			loanNumOfInstallment: 2,
			expectedDueTimeUnixes: []int64{
				time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), // New Year's Day 2024
				time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC).Unix(), // January 8, 2024
			},
		},
		{
			name:                 "Leap year February",
			startTime:            time.Date(2024, 2, 26, 10, 0, 0, 0, time.UTC), // 2024 is a leap year
			loanNumOfInstallment: 2,
			expectedDueTimeUnixes: []int64{
				time.Date(2024, 3, 4, 0, 0, 0, 0, time.UTC).Unix(),  // March 4 (after Feb 29)
				time.Date(2024, 3, 11, 0, 0, 0, 0, time.UTC).Unix(), // March 11
			},
		},
		{
			name:                 "Large number of installments",
			startTime:            time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			loanNumOfInstallment: 5,
			expectedDueTimeUnixes: []int64{
				time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC).Unix(),  // Week 1
				time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Unix(), // Week 2
				time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC).Unix(), // Week 3
				time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC).Unix(), // Week 4
				time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC).Unix(),  // Week 5
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getInstallmentDueTimeUnixes(tt.startTime, tt.loanNumOfInstallment)

			// Check if the result matches expected
			if !reflect.DeepEqual(result, tt.expectedDueTimeUnixes) {
				t.Errorf("getInstallmentDueTimeUnixes() = %v, want %v", result, tt.expectedDueTimeUnixes)

				// Provide more detailed debugging info
				t.Logf("Start time: %v", tt.startTime)
				t.Logf("Number of installments: %d", tt.loanNumOfInstallment)
				t.Logf("Expected length: %d, Got length: %d", len(tt.expectedDueTimeUnixes), len(result))

				// Convert unix timestamps back to readable dates for debugging
				if len(result) > 0 {
					t.Log("Actual due times:")
					for i, unix := range result {
						t.Logf("  [%d] %v (unix: %d)", i, time.Unix(unix, 0).UTC(), unix)
					}
				}

				if len(tt.expectedDueTimeUnixes) > 0 {
					t.Log("Expected due times:")
					for i, unix := range tt.expectedDueTimeUnixes {
						t.Logf("  [%d] %v (unix: %d)", i, time.Unix(unix, 0).UTC(), unix)
					}
				}
			}
		})
	}
}
