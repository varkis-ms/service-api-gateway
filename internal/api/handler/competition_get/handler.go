package competition_get

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

func (h *Handler) Handle(ctx context.Context, userID, competitionID int64) (model.CompetitionInfo, error) {
	competitionInfo, err := h.competitionClient.GetCompetitionInfo(ctx, userID, competitionID)
	if err != nil {
		h.log.Error("competitionClient.GetCompetitionInfo failed", sl.Err(err))
		return model.CompetitionInfo{}, err
	}

	return *competitionInfo, nil
}
