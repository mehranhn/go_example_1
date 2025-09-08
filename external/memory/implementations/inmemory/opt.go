package memoryimpinmemory

import (
	"time"
)

func (inMemory *InMemory) GetAndDeletePhoneOtpInMemory(phone string) (uint, error) {
	inMemory.mutexOtp.Lock()
	defer inMemory.mutexOtp.Unlock()

	value, exists := inMemory.mapOtp[phone]

	if !exists || time.Now().After(value.expireAt) {
		return 0, nil
	} else {
		val := value.value

		delete(inMemory.mapOtp, phone)

		return val, nil
	}
}

func (inMemory *InMemory) SetPhoneOtpInMemory(phone string, code uint, ttl time.Duration) error {
	inMemory.mutexOtp.Lock()
	defer inMemory.mutexOtp.Unlock()

	// we should run a cron job that deletes expired values after some time to free up memeory but i'm going to skip it
	inMemory.mapOtp[phone] = inMemoryValue{
		value: code,
		expireAt: time.Now().Add(ttl),
	}

	return nil
}
