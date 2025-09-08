// Package memoryinterfaces
package memoryinterfaces

import "time"

type Otp interface {
    GetAndDeletePhoneOtpInMemory(phone string) (uint, error)
    SetPhoneOtpInMemory(phone string, code uint, ttl time.Duration) error
}
