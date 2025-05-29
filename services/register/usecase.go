package register

import "context"

type RegisterUsecase interface {
	RegisterUser(ctx context.Context, params map[string]any, source string) error
}
