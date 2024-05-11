package login

import (
	"context"
)

type AuthClient interface {
	Login(ctx context.Context, email, password string) (string, error)
}
