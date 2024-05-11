package upload_file

import (
	"context"
	"log/slog"
	"mime/multipart"

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

func (h *Handler) Handle(ctx context.Context, multipartFile *multipart.FileHeader, userID, compID, fileType int64) error {
	err := h.solutionClient.UploadFile(
		ctx,
		multipartFile,
		userID,
		compID,
		fileType,
	)
	if err != nil {
		h.log.Error("solutionClient.UploadFile failed", sl.Err(err))
		return err
	}

	return nil
}
