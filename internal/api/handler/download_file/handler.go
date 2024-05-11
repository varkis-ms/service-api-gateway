package download_file

import (
	"bytes"
	"context"
	"log/slog"

	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type Handler struct {
	solutionClient SolutionClient
	log            *slog.Logger
}

func New(
	solutionClient SolutionClient,
	log *slog.Logger,
) *Handler {
	return &Handler{
		solutionClient: solutionClient,
		log:            log,
	}
}

func (h *Handler) Handle(ctx context.Context, compID int64) (*bytes.Buffer, error) {
	file, err := h.solutionClient.Download(
		ctx,
		compID,
	)
	if err != nil {
		h.log.Error("solutionClient.Download failed", sl.Err(err))
		return nil, err
	}

	return file, nil
}
