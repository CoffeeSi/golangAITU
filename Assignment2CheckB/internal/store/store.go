package store

import (
	"sync"
)

type CounterStore[K comparable] struct {
	Mu sync.Mutex
	m  map[K]int
}

func NewCounterStore[K comparable]() *CounterStore[K] {
	return &CounterStore[K]{
		m: make(map[K]int),
	}
}

func (cs *CounterStore[K]) Inc(key K) {
	cs.m[key]++
}

func (cs *CounterStore[K]) GetAll() map[K]int {
	return cs.m
}

func (cs *CounterStore[K]) Size() int {
	return len(cs.m)
}
