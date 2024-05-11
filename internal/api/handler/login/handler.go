package login

import (
	"context"
	"log/slog"

	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type Handler struct {
	authClient AuthClient
	log        *slog.Logger
}

func New(
	authClient AuthClient,
	log *slog.Logger,
) *Handler {
	return &Handler{
		authClient: authClient,
		log:        log,
	}
}

func (h *Handler) Handle(ctx context.Context, req *model.SignUpRequest) (string, error) {
	loginResponse, err := h.authClient.Login(ctx, req.Email, req.Password)
	if err != nil {
		h.log.Error("authClient.Login failed", sl.Err(err))
		return "", err
	}

	return loginResponse, nil
}
