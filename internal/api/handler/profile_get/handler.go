package profile_get

import (
	"context"
	"log/slog"

	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type Handler struct {
	userInfoClient    UserInfoClient
	competitionClient CompetitionClient
	log               *slog.Logger
}

func New(
	userInfoClient UserInfoClient,
	competitionClient CompetitionClient,
	log *slog.Logger,
) *Handler {
	return &Handler{
		userInfoClient:    userInfoClient,
		competitionClient: competitionClient,
		log:               log,
	}
}

func (h *Handler) Handle(ctx context.Context, req int64) (model.ProfileGetResponse, error) {
	out := model.ProfileGetResponse{}
	profileInfo, err := h.userInfoClient.UserInfoGet(ctx, req)
	if err != nil {
		h.log.Error("userInfoClient.UserInfoGet failed", sl.Err(err))
		return out, err
	}

	userActivityTotalInfo, err := h.competitionClient.UserActivityTotal(ctx, req)
	if err != nil {
		h.log.Error("competitionClient.UserActivityTotal failed", sl.Err(err))
		return out, err
	}

	userActivityFullInfo, err := h.competitionClient.UserActivityFull(ctx, req)
	if err != nil {
		h.log.Error("competitionClient.UserActivityFull failed", sl.Err(err))
		return out, err
	}

	out.UserMainInfo = *profileInfo
	out.UserActivityFullInfo = *userActivityFullInfo
	out.UserActivityTotalInfo = *userActivityTotalInfo
	return out, nil
}
