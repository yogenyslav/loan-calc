// Package shared переменные, типы и ошибки, которые используются в других пакетах.
package shared

import "errors"

var (
	// ErrMinInitialPayment ошибка при первоначальном взносе ниже 20% от стоимости объекта.
	ErrMinInitialPayment = errors.New("the initial payment should be more")

	// ErrMaxInitialPayment ошибка при указании первоначального взноса равным или большим общего объема кредита.
	ErrMaxInitialPayment = errors.New("the initial payment should be less")

	// ErrMultiplePrograms ошибка при выборе нескольких программ кредитования.
	ErrMultiplePrograms = errors.New("choose only 1 program")

	// ErrNoProgram ошибка при отсутствии выбора хотя бы одной из программ кредитования.
	ErrNoProgram = errors.New("choose program")

	// ErrEmptyCache ошибка при пустом кеше.
	ErrEmptyCache = errors.New("empty cache")
)
