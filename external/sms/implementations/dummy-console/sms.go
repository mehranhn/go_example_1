package smsimpdummyconsole

import "fmt"

func (*DummyConsole) SendSms(_ string, message string) error {
	fmt.Println(message)

	return nil
}
