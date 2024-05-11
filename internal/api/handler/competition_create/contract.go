package competition_create

import (
	"context"

	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type CompetitionClient interface {
	CompetitionCreate(ctx context.Context, compInfo *model.CompetitionCreate) (int64, error)
}
