package save_solution

import (
	"context"
)

type CompetitionClient interface {
	SaveSolution(ctx context.Context, userID, competitionID int64) error
}
