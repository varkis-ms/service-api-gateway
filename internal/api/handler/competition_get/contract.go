package competition_get

import (
	"context"

	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type CompetitionClient interface {
	GetCompetitionInfo(ctx context.Context, userID, competitionID int64) (*model.CompetitionInfo, error)
}
