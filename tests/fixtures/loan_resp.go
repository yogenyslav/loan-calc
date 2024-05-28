package fixtures

import (
	"loan/internal/loan_calc/model"
	"loan/internal/shared"
	"time"
)

// LoanRespBuilder builder для model.LoanResp.
type LoanRespBuilder struct {
	LoanResp *model.LoanResp
}

// LoanResp конструктор для LoanRespBuilder.
func LoanResp() *LoanRespBuilder {
	return &LoanRespBuilder{
		LoanResp: &model.LoanResp{},
	}
}

// New создать новый model.LoanResp по умолчанию.
func (b *LoanRespBuilder) New() *LoanRespBuilder {
	t := time.Now().AddDate(0, 240, 0)

	return b.
		LastPaymentDate(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)).
		Rate(shared.RateSalary).
		LoanSum(4000000).
		MonthlyPayment(33458).
		Overpayment(4029920).
		ObjectCost(5000000).
		InitialPayment(1000000).
		Months(240)
}

// V получить значение model.LoanResp из LoanRespBuilder.
func (b *LoanRespBuilder) V() model.LoanResp {
	return *b.LoanResp
}

// P получить указатель на model.LoanResp из LoanRespBuilder.
func (b *LoanRespBuilder) P() *model.LoanResp {
	return b.LoanResp
}

// Aggregates задать значение поля Aggregates.
func (b *LoanRespBuilder) Aggregates(v model.Aggregates) *LoanRespBuilder {
	b.LoanResp.Aggregates = v
	return b
}

// LastPaymentDate задать значение поля LastPaymentDate.
func (b *LoanRespBuilder) LastPaymentDate(v time.Time) *LoanRespBuilder {
	b.LoanResp.Aggregates.LastPaymentDate = v
	return b
}

// Rate задать значение поля Rate.
func (b *LoanRespBuilder) Rate(v shared.ProgramRate) *LoanRespBuilder {
	b.LoanResp.Aggregates.Rate = v
	return b
}

// LoanSum задать значение поля LoanSum.
func (b *LoanRespBuilder) LoanSum(v int) *LoanRespBuilder {
	b.LoanResp.Aggregates.LoanSum = v
	return b
}

// MonthlyPayment задать значение поля MonthlyPayment.
func (b *LoanRespBuilder) MonthlyPayment(v int) *LoanRespBuilder {
	b.LoanResp.Aggregates.MonthlyPayment = v
	return b
}

// Overpayment задать значение поля Overpayment.
func (b *LoanRespBuilder) Overpayment(v int) *LoanRespBuilder {
	b.LoanResp.Aggregates.Overpayment = v
	return b
}

// Params задать значение поля Params.
func (b *LoanRespBuilder) Params(v model.Params) *LoanRespBuilder {
	b.LoanResp.Params = v
	return b
}

// ObjectCost задать значение поля ObjectCost.
func (b *LoanRespBuilder) ObjectCost(v int) *LoanRespBuilder {
	b.LoanResp.Params.ObjectCost = v
	return b
}

// InitialPayment задать значение поля InitialPayment.
func (b *LoanRespBuilder) InitialPayment(v int) *LoanRespBuilder {
	b.LoanResp.Params.InitialPayment = v
	return b
}

// Months задать значение поля Months.
func (b *LoanRespBuilder) Months(v int) *LoanRespBuilder {
	b.LoanResp.Params.Months = v
	return b
}

// Program задать значение поля Program.
func (b *LoanRespBuilder) Program(v model.Program) *LoanRespBuilder {
	b.LoanResp.Program = v
	return b
}

// Salary задать значение Salary поля Program.
func (b *LoanRespBuilder) Salary(v bool) *LoanRespBuilder {
	b.LoanResp.Program.Salary = v
	return b
}

// Military задать значение Military поля Program.
func (b *LoanRespBuilder) Military(v bool) *LoanRespBuilder {
	b.LoanResp.Program.Military = v
	return b
}

// Base задать значение Base поля Program.
func (b *LoanRespBuilder) Base(v bool) *LoanRespBuilder {
	b.LoanResp.Program.Base = v
	return b
}
