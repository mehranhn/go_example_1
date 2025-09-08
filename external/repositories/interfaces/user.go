package repointerfaces

import (
	"github.com/google/uuid"
	"github.com/mehranhn/go_example_1/models/request"
	"github.com/mehranhn/go_example_1/models/response"
)

type User interface {
    GetUserById(id uuid.UUID) (*response.UserDto, error)
    GetUserByPhone(phone string) (*response.UserDto, error)
    GetUserList(filter request.PaginationFilter) ([]response.UserDto, error)
}
