// Package utils
package utils

import "math/rand"

func GenerateOtp() uint {
	min := 10000
	max := 100000

	return uint(rand.Intn(max-min) + min)
}
