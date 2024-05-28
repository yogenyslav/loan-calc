package shared

import "time"

const (
	// InitialPaymentThreshold относительное пороговое значение минимального первоначального взноса.
	InitialPaymentThreshold float64 = 0.2
	// CacheExp время до инвалидации кеша.
	CacheExp = 24 * time.Hour
)

// ProgramRate процентная ставка в зависимости от программы кредитования.
type ProgramRate int8

const (
	_ ProgramRate = iota
	// RateSalary ставка для корпоративных клиентов.
	RateSalary = 8
	// RateMilitary военная ипотека.
	RateMilitary = 9
	// RateBase базовая программа.
	RateBase = 10
)
