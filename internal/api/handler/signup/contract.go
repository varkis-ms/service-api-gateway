package signup

import (
	"context"
)

type AuthClient interface {
	Signup(ctx context.Context, email, password string) (int64, error)
}

type UserInfoClient interface {
	UserInfoSave(ctx context.Context, userID int64, email string) error
}
