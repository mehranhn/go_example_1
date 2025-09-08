package repositoryimppostgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/mehranhn/go_example_1/models/entities"
	"github.com/mehranhn/go_example_1/models/request"
	"github.com/mehranhn/go_example_1/models/response"
)

func (pg *Postgres) GetUserById(id uuid.UUID) (*response.UserDto, error) {
	var user entities.UserEntity
	err := pg.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	userDto := user.ToDto()
	return &userDto, nil
}

func (pg *Postgres) GetUserByPhone(phone string) (*response.UserDto, error) {
	var user entities.UserEntity
	err := pg.db.Get(&user, "SELECT * FROM users WHERE phone = $1", phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	userDto := user.ToDto()
	return &userDto, nil
}

func (pg *Postgres) GetUserList(filter request.PaginationFilter) ([]response.UserDto, error) {
	users := make([]entities.UserEntity, 0, filter.Limit)
	args := []any{}
	argCounter := 1

	query := `SELECT * FROM users`

	if filter.Search != "" {
		query += " WHERE phone LIKE $1"
		args = append(args, "%"+filter.Search+"%")
		argCounter += 1
	}

	query += fmt.Sprintf(" LIMIT $%d", argCounter)
	argCounter += 1
	args = append(args, filter.Limit)
	query += fmt.Sprintf(" OFFSET $%d", argCounter)
	args = append(args, filter.Offset())
	argCounter += 1

	err := pg.db.Select(&users, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return []response.UserDto{}, nil
		}

		return []response.UserDto{}, err
	}

	userDtos := make([]response.UserDto, 0, len(users))
	for _, u := range users {
		userDtos = append(userDtos, u.ToDto())
	}

	return userDtos, nil
}
