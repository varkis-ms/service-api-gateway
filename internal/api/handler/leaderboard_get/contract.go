package leaderboard_get

import (
	"context"

	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type CompetitionClient interface {
	LeaderBoard(ctx context.Context, competitionID int64) (*model.LeaderboardList, error)
}
