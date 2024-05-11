package competition_list_get

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

func (h *Handler) Handle(ctx context.Context, userID int64) ([]model.CompetitionInfoSmall, error) {
	competitionList, err := h.competitionClient.CompetitionList(ctx, userID)
	if err != nil {
		h.log.Error("competitionClient.CompetitionList failed", sl.Err(err))
		return nil, err
	}

	return competitionList, nil
}
