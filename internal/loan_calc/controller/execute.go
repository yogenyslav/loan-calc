package controller

import (
	"context"
	"loan/internal/loan_calc/model"
	"loan/internal/shared"
	"math"
	"time"
)

// Execute рассчитать параметры кредитования.
func (ctrl *Controller) Execute(_ context.Context, params model.LoanReq) (model.LoanResp, error) {
	var resp model.LoanResp

	resp.Params.ObjectCost = params.ObjectCost
	resp.Params.InitialPayment = params.InitialPayment
	resp.Params.Months = params.Months
	resp.Program = params.Program

	if float64(params.InitialPayment) < float64(params.ObjectCost)*shared.InitialPaymentThreshold {
		return resp, shared.ErrMinInitialPayment
	}

	if params.InitialPayment >= params.ObjectCost {
		return resp, shared.ErrMaxInitialPayment
	}

	programChoices := []bool{params.Program.Salary, params.Program.Military, params.Program.Base}
	flag := false
	for _, choice := range programChoices {
		if choice {
			if flag {
				return resp, shared.ErrMultiplePrograms
			}
			flag = true
		}
	}
	if !flag {
		return resp, shared.ErrNoProgram
	}

	resp.Aggregates.Rate = getProgramRate(params.Program)
	resp.Aggregates.LoanSum = params.ObjectCost - params.InitialPayment

	resp.Aggregates.MonthlyPayment = calcMonthlyPayment(resp.Aggregates.Rate, params.Months, resp.Aggregates.LoanSum)
	resp.Aggregates.Overpayment = resp.Aggregates.MonthlyPayment*params.Months - resp.Aggregates.LoanSum

	t := time.Now().AddDate(0, params.Months, 0)
	resp.Aggregates.LastPaymentDate = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	ctrl.cache.Insert(resp, shared.CacheExp)

	return resp, nil
}

func getProgramRate(program model.Program) shared.ProgramRate {
	if program.Base {
		return shared.RateBase
	}
	if program.Military {
		return shared.RateMilitary
	}
	if program.Salary {
		return shared.RateSalary
	}
	return 0
}

func calcMonthlyPayment(rate shared.ProgramRate, months, loanSum int) int {
	monthlyRate := float64(rate) / 12 / 100
	coef := (monthlyRate * math.Pow(1+monthlyRate, float64(months))) / (math.Pow(1+monthlyRate, float64(months)) - 1)
	return int(math.Ceil(float64(loanSum) * coef))
}
