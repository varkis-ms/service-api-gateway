package mwauth

import (
	"context"
)

type AuthClient interface {
	Validate(ctx context.Context, token string) (int64, error)
}
