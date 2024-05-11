package profile_get

import (
	"context"

	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type UserInfoClient interface {
	UserInfoGet(ctx context.Context, userID int64) (*model.UserInfo, error)
}

type CompetitionClient interface {
	UserActivityTotal(ctx context.Context, userID int64) (*model.UserActivityTotal, error)
	UserActivityFull(ctx context.Context, userID int64) (*model.UserActivityFull, error)
}
