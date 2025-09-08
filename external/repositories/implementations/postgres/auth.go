package repositoryimppostgres

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mehranhn/go_example_1/constants"
	"github.com/mehranhn/go_example_1/models/request"
)

func (pg *Postgres) UpsertUser(data request.RegisterOrLoginDto) (constants.RegisterOrLoginResult, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return constants.Login, err
	}

	_, err = pg.db.Exec("INSERT INTO users (id, phone) VALUES ($1, $2)", newId, data.Phone)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				if pqErr.Constraint == "users_phone_key" {
					return constants.Login, nil
				}
			}
		}

		return constants.Login, err
	}

	return constants.Register, nil
}
