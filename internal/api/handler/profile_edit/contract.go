package profile_edit

import (
	"context"

	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type UserInfoClient interface {
	UserInfoEdit(ctx context.Context, userInfo *model.UserInfo) error
}
