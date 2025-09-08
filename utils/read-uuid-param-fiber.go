package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ReadUUIDParamFiber(c *fiber.Ctx, paramName string) (uuid.UUID, error) {
	idStr := c.Params(paramName)

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
