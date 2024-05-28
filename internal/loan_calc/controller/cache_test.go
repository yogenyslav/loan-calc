package controller_test

import (
	"context"
	"loan/internal/loan_calc/controller"
	"loan/internal/loan_calc/model"
	"loan/internal/shared"
	"loan/pkg/storage"
	"loan/tests/fixtures"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestController_List(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		elems     []model.LoanReq
		wantElems []model.CacheResp
		wantErr   error
	}{
		{
			name:  "success, 1 element",
			elems: []model.LoanReq{fixtures.LoanReq().New().Salary(true).V()},
			wantElems: []model.CacheResp{
				{LoanResp: fixtures.LoanResp().New().Salary(true).V(), ID: 0},
			},
			wantErr: nil,
		},
		{
			name: "success, many elements",
			elems: []model.LoanReq{
				fixtures.LoanReq().New().Salary(true).V(),
				fixtures.LoanReq().New().Salary(true).V(),
				fixtures.LoanReq().New().Salary(true).V(),
			},
			wantElems: []model.CacheResp{
				{LoanResp: fixtures.LoanResp().New().Salary(true).V(), ID: 0},
				{LoanResp: fixtures.LoanResp().New().Salary(true).V(), ID: 1},
				{LoanResp: fixtures.LoanResp().New().Salary(true).V(), ID: 2},
			},
			wantErr: nil,
		},
		{
			name:    "fail, empty cache",
			wantErr: shared.ErrEmptyCache,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				err   error
				ctx   = context.Background()
				cache = storage.NewInMemCache[model.LoanResp]()
				ctrl  = controller.New(cache)
			)

			for _, elem := range tt.elems {
				_, err = ctrl.Execute(ctx, elem)
				require.NoError(t, err)
			}

			records, err := ctrl.List(ctx)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.wantElems, records)
			}
		})
	}
}
