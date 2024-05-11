package competition_create

import (
	"context"
	"log/slog"

	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type Handler struct {
	competitionClient CompetitionClient
	log               *slog.Logger
}

func New(
	competitionClient CompetitionClient,
	log *slog.Logger,
) *Handler {
	return &Handler{
		competitionClient: competitionClient,
		log:               log,
	}
}

func (h *Handler) Handle(ctx context.Context, req *model.CompetitionCreate) (int64, error) {
	competitionID, err := h.competitionClient.CompetitionCreate(ctx, req)
	if err != nil {
		h.log.Error("competitionClient.CompetitionCreate failed", sl.Err(err))
		return 0, err
	}

	return competitionID, nil
}
