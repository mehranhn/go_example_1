// Package smsinterfaces
package smsinterfaces

type Sms interface {
    SendSms(phone string, message string) error
}
