package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/mehranhn/go_example_1/models/response"
)

type UserEntity struct {
	ID        uuid.UUID `db:"id"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

func (user *UserEntity) ToDto() response.UserDto {
	return response.UserDto{
		ID: user.ID,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
	}
}
