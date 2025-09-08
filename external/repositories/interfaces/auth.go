// Package repointerfaces
package repointerfaces

import (
	"github.com/mehranhn/go_example_1/constants"
	"github.com/mehranhn/go_example_1/models/request"
)

type Auth interface {
	UpsertUser(data request.RegisterOrLoginDto) (constants.RegisterOrLoginResult, error)
}
