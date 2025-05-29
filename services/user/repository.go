package user

import (
	"github.com/gofrs/uuid"
)

type UserRepository interface {
	RegisterUser(params map[string]any) (*uuid.UUID, error)
}
