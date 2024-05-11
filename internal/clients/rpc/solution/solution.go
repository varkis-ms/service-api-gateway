package authrpc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/varkis-ms/service-api-gateway/internal/pkg/pb/clients"
	"google.golang.org/grpc"
)

const chunkSize = 1024 * 1024

type SolutionClient struct {
	api clients.SolutionClient
}

func New(
	rpcClient *grpc.ClientConn,
) *SolutionClient {
	return &SolutionClient{
		api: clients.NewSolutionClient(rpcClient),
	}
}

func (c *SolutionClient) UploadFile(ctx context.Context, multipartFile *multipart.FileHeader, userID, competitionID, fileType int64) error {
	file, err := multipartFile.Open()
	if err != nil {
		return errors.New("invalid file")
	}

	defer file.Close()

	stream, err := c.api.SaveFile(ctx)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	buf := make([]byte, chunkSize)
	batchNumber := 1
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := buf[:num]

		if err := stream.Send(&clients.FileUploadRequest{
			Chunk:         chunk,
			UserID:        userID,
			CompetitionID: competitionID,
			Type:          clients.Type(int32(fileType)),
		}); err != nil {
			return err
		}
		batchNumber += 1
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}

func (c *SolutionClient) Download(ctx context.Context, competitionID int64) (*bytes.Buffer, error) {
	stream, err := c.api.Download(
		ctx,
		&clients.DownloadRequest{CompetitionID: competitionID},
	)
	if err != nil {
		return nil, fmt.Errorf("SolutionClient.Download: %w", err)
	}

	file := new(bytes.Buffer)
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		chunk := req.GetChunk()

		if _, err = file.Write(chunk); err != nil {
			return nil, err
		}
	}

	return file, nil
}
