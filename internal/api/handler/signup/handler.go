package signup

import (
	"context"
	"log/slog"

	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type Handler struct {
	authClient     AuthClient
	userInfoClient UserInfoClient
	log            *slog.Logger
}

func New(
	authClient AuthClient,
	userInfoClient UserInfoClient,
	log *slog.Logger,
) *Handler {
	return &Handler{
		authClient:     authClient,
		userInfoClient: userInfoClient,
		log:            log,
	}
}

func (h *Handler) Handle(ctx context.Context, req *model.SignUpRequest) error {
	userID, err := h.authClient.Signup(ctx, req.Email, req.Password)
	if err != nil {
		h.log.Error("authClient.Signup failed", sl.Err(err))
		return err
	}

	if err = h.userInfoClient.UserInfoSave(ctx, userID, req.Email); err != nil {
		h.log.Error("userInfoClient.UserInfoSave failed", sl.Err(err))
		return err
	}

	return nil
}
