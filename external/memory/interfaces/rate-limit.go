package memoryinterfaces

import "time"

type RateLimit interface {
	FetchAddKey(key string, ttl time.Duration) (uint, error)
}
