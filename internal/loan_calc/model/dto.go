// Package model хранит описание структур всех запросов и ответов сервиса.
package model

import (
	"loan/internal/shared"
	"time"
)

// LoanReq модель запроса с параметрами кредита и выбором программы.
type LoanReq struct {
	// ObjectCost     int     `json:"object_cost"`
	// InitialPayment int     `json:"initial_payment"`
	// Months         int     `json:"months"`
	Params
	Program Program `json:"program"`
}

// Program опции программы кредитования.
type Program struct {
	Salary   bool `json:"salary"`
	Military bool `json:"military"`
	Base     bool `json:"base"`
}

// LoanResp модель ответа со входными и рассчитанными параметрами.
type LoanResp struct {
	Aggregates Aggregates `json:"aggregates"`
	Params     Params     `json:"params"`
	Program    Program    `json:"program"`
}

// Params основные параметры кредита.
type Params struct {
	ObjectCost     int `json:"object_cost"`
	InitialPayment int `json:"initial_payment"`
	Months         int `json:"months"`
}

// Aggregates рассчитываемые праметры кредитования.
type Aggregates struct {
	LastPaymentDate time.Time          `json:"last_payment_date"`
	Rate            shared.ProgramRate `json:"rate"`
	LoanSum         int                `json:"loan_sum"`
	MonthlyPayment  int                `json:"monthly_payment"`
	Overpayment     int                `json:"overpayment"`
}

// CacheResp ответ на запрос чтения данных из кеша.
type CacheResp struct {
	LoanResp
	ID int `json:"id"`
}
