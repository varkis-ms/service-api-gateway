package leaderboard_get

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

func (h *Handler) Handle(ctx context.Context, competitionID int64) (model.LeaderboardList, error) {
	leaderboardList, err := h.competitionClient.LeaderBoard(ctx, competitionID)
	if err != nil {
		h.log.Error("competitionClient.GetCompetitionInfo failed", sl.Err(err))
		return model.LeaderboardList{}, err
	}

	return *leaderboardList, nil
}
