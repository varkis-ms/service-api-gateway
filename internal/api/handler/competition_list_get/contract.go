package competition_list_get

import (
	"context"

	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type CompetitionClient interface {
	CompetitionList(ctx context.Context, userID int64) ([]model.CompetitionInfoSmall, error)
}
