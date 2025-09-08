// Package memoryimpinmemory stores the application state in the process memory
package memoryimpinmemory

import (
	"sync"
	"time"
)

type inMemoryValue struct {
	value    uint
	expireAt time.Time
}

type InMemory struct {
	mutexRateLimit       sync.Mutex
	mutexOtp             sync.Mutex
	mapRateLimit         map[string]inMemoryValue
	mapOtp               map[string]inMemoryValue
}

func NewInMemory() InMemory {
	return InMemory{
		mutexRateLimit:       sync.Mutex{},
		mutexOtp:             sync.Mutex{},
		mapRateLimit:         make(map[string]inMemoryValue),
		mapOtp:               make(map[string]inMemoryValue),
	}
}
