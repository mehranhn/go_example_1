package memoryimpinmemory

import (
	"time"
)

func (inMemory *InMemory) FetchAddKey(key string, ttl time.Duration) (uint, error) {
	inMemory.mutexRateLimit.Lock()
	defer inMemory.mutexRateLimit.Unlock()

	value, exists := inMemory.mapRateLimit[key]

	if !exists || time.Now().After(value.expireAt) {
		inMemory.mapRateLimit[key] = inMemoryValue{
			value: 2,
			expireAt: time.Now().Add(ttl),
		}

		return 1, nil;
	} else {
		inMemory.mapRateLimit[key] = inMemoryValue{
			value: value.value + 1,
			expireAt: value.expireAt,
		}

		return value.value, nil;
	}
}
