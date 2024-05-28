package controller_test

import (
	"context"
	"loan/internal/loan_calc/controller"
	"loan/internal/loan_calc/model"
	"loan/internal/shared"
	"loan/pkg/storage"
	"loan/tests/fixtures"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestController_Execute(t *testing.T) {
	t.Parallel()

	cache := storage.NewInMemCache[model.LoanResp]()
	ctrl := controller.New(cache)

	tests := []struct {
		name     string
		params   model.LoanReq
		wantResp model.LoanResp
		wantErr  error
	}{
		{
			name:     "success, program salary",
			params:   fixtures.LoanReq().New().Salary(true).V(),
			wantResp: fixtures.LoanResp().New().Salary(true).V(),
			wantErr:  nil,
		},
		{
			name:     "success, program military",
			params:   fixtures.LoanReq().New().Military(true).V(),
			wantResp: fixtures.LoanResp().New().MonthlyPayment(35990).Overpayment(4637600).LastPaymentDate(time.Date(2044, 5, 28, 0, 0, 0, 0, time.UTC)).Rate(shared.RateMilitary).Military(true).V(),
			wantErr:  nil,
		},
		{
			name:     "success, program base",
			params:   fixtures.LoanReq().New().Base(true).V(),
			wantResp: fixtures.LoanResp().New().MonthlyPayment(38601).Overpayment(5264240).LastPaymentDate(time.Date(2044, 5, 28, 0, 0, 0, 0, time.UTC)).Rate(shared.RateBase).Base(true).V(),
			wantErr:  nil,
		},
		{
			name:    "fail, no program",
			params:  fixtures.LoanReq().ObjectCost(10).InitialPayment(8).V(),
			wantErr: shared.ErrNoProgram,
		},
		{
			name:    "fail, more than one program",
			params:  fixtures.LoanReq().Base(true).ObjectCost(10).InitialPayment(8).Salary(true).V(),
			wantErr: shared.ErrMultiplePrograms,
		},
		{
			name:    "fail, initial should be more",
			params:  fixtures.LoanReq().Base(true).ObjectCost(1000000).InitialPayment(1).V(),
			wantErr: shared.ErrMinInitialPayment,
		},
		{
			name:    "fail, initial payment exceedes object cost",
			params:  fixtures.LoanReq().Base(true).ObjectCost(1000000).InitialPayment(1000001).V(),
			wantErr: shared.ErrMaxInitialPayment,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resp, err := ctrl.Execute(context.Background(), tt.params)
			if err != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.wantResp, resp)
			}
		})
	}
}
