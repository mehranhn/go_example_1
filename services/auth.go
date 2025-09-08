// Package service
package services

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mehranhn/go_example_1/constants"
	memoryinterfaces "github.com/mehranhn/go_example_1/external/memory/interfaces"
	repointerfaces "github.com/mehranhn/go_example_1/external/repositories/interfaces"
	smsinterfaces "github.com/mehranhn/go_example_1/external/sms/interfaces"
	"github.com/mehranhn/go_example_1/models/request"
	"github.com/mehranhn/go_example_1/models/response"
	"github.com/mehranhn/go_example_1/utils"
)

type memoryinterface interface {
	memoryinterfaces.RateLimit
	memoryinterfaces.Otp
}

type repoInterface interface {
	repointerfaces.Auth
	repointerfaces.User
}

type AuthService struct {
	repo                 repoInterface
	memory               memoryinterface
	sms                  smsinterfaces.Sms
	maxRateLimitAttempts uint
	rateLimitTTLDuration time.Duration
	otpTTLDuration       time.Duration
	jwtSecretKey         string
	jwtTTLDuration       time.Duration
}

func NewAuthService(authRepo repoInterface, memory memoryinterface, sms smsinterfaces.Sms, maxRateLimitAttempts uint, rateLimitTTLDuration time.Duration, otpTTLDuration time.Duration, jwtSecretKey string, jwtTTLDuration time.Duration) AuthService {
	return AuthService{
		authRepo,
		memory,
		sms,
		maxRateLimitAttempts,
		rateLimitTTLDuration,
		otpTTLDuration,
		jwtSecretKey,
		jwtTTLDuration,
	}
}

func (service *AuthService) RegistryOrLogin(data request.RegisterOrLoginDto) (*constants.RegisterOrLoginResult, error) {
	attempts, err := service.memory.FetchAddKey(data.Phone, service.rateLimitTTLDuration)
	if err != nil {
		return nil, err
	}

	if service.maxRateLimitAttempts < attempts {
		return nil, fiber.NewError(fiber.StatusTooManyRequests)
	}

	res, err := service.repo.UpsertUser(data)
	if err != nil {
		return nil, err
	}

	otpCode := utils.GenerateOtp()

	err = service.memory.SetPhoneOtpInMemory(data.Phone, otpCode, service.otpTTLDuration)
	if err != nil {
		return nil, err
	}

	err = service.sms.SendSms(data.Phone, fmt.Sprintf("otp code: %d", otpCode))
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ConfirmOtp godoc
// return code 0 success
// return code 1 wrong otp
// return code 2 err
func (service *AuthService) ConfirmOtp(data request.ConfirmOtpDto) (uint8, response.TokenDto, error) {
	value, err := service.memory.GetAndDeletePhoneOtpInMemory(data.Phone)
	if err != nil {
		return 2, response.TokenDto{Token: ""}, err
	}

	if value != data.Code {
		return 1, response.TokenDto{Token: ""}, nil
	}

	user, err := service.repo.GetUserByPhone(data.Phone)
	if err != nil {
		return 2, response.TokenDto{Token: ""}, err
	}

	token, err := utils.GenerateJWTToken(service.jwtSecretKey, user.ID.String(), user.Phone, service.jwtTTLDuration)
	if err != nil {
		return 2, response.TokenDto{Token: ""}, err
	}

	return 0, response.TokenDto{Token: token}, nil
}
