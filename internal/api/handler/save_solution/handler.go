package save_solution

import (
	"context"
	"log/slog"

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

func (h *Handler) Handle(ctx context.Context, userID, competitionID int64) error {
	err := h.competitionClient.SaveSolution(ctx, userID, competitionID)
	if err != nil {
		h.log.Error("competitionClient.SaveSolution failed", sl.Err(err))
		return err
	}

	return nil
}
