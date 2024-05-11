package upload_file

import (
	"context"
	"mime/multipart"
)

type SolutionClient interface {
	UploadFile(ctx context.Context, multipartFile *multipart.FileHeader, userID, competitionID, fileType int64) error
}
