package competition_edit

import (
	"context"

	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type CompetitionClient interface {
	CompetitionEdit(ctx context.Context, compInfo *model.CompetitionEdit) error
}
