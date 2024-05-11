package profile_edit

import (
	"context"
	"log/slog"

	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type Handler struct {
	userInfoClient UserInfoClient
	log            *slog.Logger
}

func New(
	userInfoClient UserInfoClient,
	log *slog.Logger,
) *Handler {
	return &Handler{
		userInfoClient: userInfoClient,
		log:            log,
	}
}

func (h *Handler) Handle(ctx context.Context, req *model.UserInfo) error {
	if err := h.userInfoClient.UserInfoEdit(ctx, req); err != nil {
		h.log.Error("userInfoClient.UserInfoEdit failed", sl.Err(err))
		return err
	}

	return nil
}
