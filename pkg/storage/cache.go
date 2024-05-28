// Package storage реализация внутренних и внешних хранилищ/кешей.
package storage

import (
	"sync"
	"time"
)

type record[T any] struct {
	v   T
	exp time.Time
}

// InMemCache основной объект кеша, хранит все записанные в него данные.
type InMemCache[T any] struct {
	data []*record[T]
	mu   sync.Mutex
}

// NewInMemCache конструктор для InMemCache.
func NewInMemCache[T any]() *InMemCache[T] {
	return &InMemCache[T]{
		data: make([]*record[T], 0),
	}
}

// Insert добавить запись в кеш.
func (c *InMemCache[T]) Insert(v T, exp time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.tryRealloc()

	c.data = append(c.data, &record[T]{
		v:   v,
		exp: time.Now().Add(exp),
	})
}

// List получить все записи из кеша и, при необходимости, инвалидация.
func (c *InMemCache[T]) List() []T {
	c.mu.Lock()
	defer c.mu.Unlock()

	res := make([]T, len(c.data))
	p := 0

	for idx, obj := range c.data {
		if time.Since(obj.exp) >= 0 {
			c.data = append(c.data[:idx], c.data[idx+1:]...)
			continue
		}
		res[p] = obj.v
		p++
	}
	return res
}

func (c *InMemCache[T]) tryRealloc() {
	if len(c.data)*2 < cap(c.data) {
		newData := make([]*record[T], len(c.data))
		copy(newData, c.data)
		c.data = newData
	}
}
