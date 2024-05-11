package competition_edit

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

func (h *Handler) Handle(ctx context.Context, req *model.CompetitionEdit) error {
	if err := h.competitionClient.CompetitionEdit(ctx, req); err != nil {
		h.log.Error("competitionClient.CompetitionEdit failed", sl.Err(err))
		return err
	}

	return nil
}
