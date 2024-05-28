package controller

import (
	"context"
	"loan/internal/loan_calc/model"
	"loan/internal/shared"
)

// List достать данные из кеша.
func (ctrl *Controller) List(_ context.Context) ([]model.CacheResp, error) {
	records := ctrl.cache.List()
	if len(records) == 0 {
		return nil, shared.ErrEmptyCache
	}

	res := make([]model.CacheResp, len(records))
	for idx, record := range records {
		res[idx] = model.CacheResp{
			LoanResp: record,
			ID:       idx,
		}
	}

	return res, nil
}
