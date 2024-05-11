package download_file

import (
	"bytes"
	"context"
)

type SolutionClient interface {
	Download(ctx context.Context, competitionID int64) (*bytes.Buffer, error)
}
