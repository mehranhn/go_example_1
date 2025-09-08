package services

import (
	"github.com/google/uuid"
	repointerfaces "github.com/mehranhn/go_example_1/external/repositories/interfaces"
	"github.com/mehranhn/go_example_1/models/request"
	"github.com/mehranhn/go_example_1/models/response"
)

type UserService struct {
	userRepo repointerfaces.User
}

func NewUserService(userRepo repointerfaces.User) UserService {
	return UserService{
		userRepo,
	}
}

func (service *UserService) GetUserById(id uuid.UUID) (*response.UserDto, error) {
	return service.userRepo.GetUserById(id)
}

func (service *UserService) GetUserList(filter request.PaginationFilter) ([]response.UserDto, error) {
	return service.userRepo.GetUserList(filter)
}
