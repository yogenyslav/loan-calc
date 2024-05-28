// Package fixtures builder тестовых данных.
package fixtures

import (
	"loan/internal/loan_calc/model"
)

// LoanReqBuilder builder для model.LoanReq.
type LoanReqBuilder struct {
	LoanReq *model.LoanReq
}

// LoanReq конструктор для LoanReqBuilder.
func LoanReq() *LoanReqBuilder {
	return &LoanReqBuilder{
		LoanReq: &model.LoanReq{},
	}
}

// New создать новый model.LoanReq по умолчанию.
func (b *LoanReqBuilder) New() *LoanReqBuilder {
	return b.
		ObjectCost(5000000).
		InitialPayment(1000000).
		Months(240)
}

// V получить значение model.LoanReq из LoanReqBuilder.
func (b *LoanReqBuilder) V() model.LoanReq {
	return *b.LoanReq
}

// P получить указатель на model.LoanReq из LoanReqBuilder.
func (b *LoanReqBuilder) P() *model.LoanReq {
	return b.LoanReq
}

// ObjectCost задать значение поля ObjectCost.
func (b *LoanReqBuilder) ObjectCost(v int) *LoanReqBuilder {
	b.LoanReq.ObjectCost = v
	return b
}

// InitialPayment задать значение поля InitialPayment.
func (b *LoanReqBuilder) InitialPayment(v int) *LoanReqBuilder {
	b.LoanReq.InitialPayment = v
	return b
}

// Months задать значение поля Months.
func (b *LoanReqBuilder) Months(v int) *LoanReqBuilder {
	b.LoanReq.Months = v
	return b
}

// Program задать значение поля Program.
func (b *LoanReqBuilder) Program(v model.Program) *LoanReqBuilder {
	b.LoanReq.Program = v
	return b
}

// Salary задать значение Salary поля Program.
func (b *LoanReqBuilder) Salary(v bool) *LoanReqBuilder {
	b.LoanReq.Program.Salary = v
	return b
}

// Military задать значение Military поля Program.
func (b *LoanReqBuilder) Military(v bool) *LoanReqBuilder {
	b.LoanReq.Program.Military = v
	return b
}

// Base задать значение Base поля Program.
func (b *LoanReqBuilder) Base(v bool) *LoanReqBuilder {
	b.LoanReq.Program.Base = v
	return b
}
