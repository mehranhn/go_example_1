package response

import (
	"time"

	"github.com/google/uuid"
)

type UserDto struct {
	ID        uuid.UUID
	Phone     string
	CreatedAt time.Time
}
