package competitionrpc

import (
	"context"
	"fmt"

	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/pb/clients"

	"google.golang.org/grpc"
)

type CompetitionClient struct {
	api clients.CompetitionClient
}

func New(
	rpcClient *grpc.ClientConn,
) *CompetitionClient {
	return &CompetitionClient{
		api: clients.NewCompetitionClient(rpcClient),
	}
}

func (c *CompetitionClient) CompetitionCreate(ctx context.Context, compInfo *model.CompetitionCreate) (int64, error) {
	response, err := c.api.CompetitionCreate(ctx, &clients.CompetitionCreateRequest{
		UserID:             compInfo.UserID,
		Title:              compInfo.Title,
		Description:        compInfo.Description,
		DatasetTitle:       compInfo.DatasetTitle,
		DatasetDescription: compInfo.DatasetDescription,
	})
	if err != nil {
		return 0, fmt.Errorf("grpc.Competition.CompetitionCreate failed: %w", err)
	}

	return response.CompetitionID, nil
}

func (c *CompetitionClient) CompetitionEdit(ctx context.Context, compInfo *model.CompetitionEdit) error {
	_, err := c.api.CompetitionEdit(ctx, &clients.CompetitionEditRequest{
		UserID:             compInfo.UserID,
		CompetitionID:      compInfo.CompetitionID,
		Title:              compInfo.Title,
		Description:        compInfo.Description,
		DatasetTitle:       compInfo.DatasetTitle,
		DatasetDescription: compInfo.DatasetDescription,
	})
	if err != nil {
		return fmt.Errorf("grpc.Competition.CompetitionEdit failed: %w", err)
	}

	return nil
}

func (c *CompetitionClient) CompetitionList(ctx context.Context, userID int64) ([]model.CompetitionInfoSmall, error) {
	response, err := c.api.CompetitionList(ctx, &clients.CompetitionListRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc.Competition.CompetitionList failed: %w", err)
	}

	return c.competitionListResponseToDto(response), nil
}

func (c *CompetitionClient) GetCompetitionInfo(ctx context.Context, userID, competitionID int64) (*model.CompetitionInfo, error) {
	response, err := c.api.GetCompetitionInfo(ctx, &clients.CompetitionInfoRequest{
		UserID:        userID,
		CompetitionID: competitionID,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc.Competition.GetCompetitionInfo failed: %w", err)
	}

	return c.competitionInfoResponseToDto(response), nil
}

func (c *CompetitionClient) LeaderBoard(ctx context.Context, competitionID int64) (*model.LeaderboardList, error) {
	response, err := c.api.LeaderBoard(ctx, &clients.LeaderBoardRequest{
		CompetitionID: competitionID,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc.Competition.LeaderBoard failed: %w", err)
	}

	return c.leaderboardResponseToDto(response), nil
}

func (c *CompetitionClient) UserActivityTotal(ctx context.Context, userID int64) (*model.UserActivityTotal, error) {
	response, err := c.api.UserActivityTotal(ctx, &clients.UserActivityTotalRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc.Competition.UserActivityTotal failed: %w", err)
	}

	return c.userActivityTotalResponseToDto(response), nil
}

func (c *CompetitionClient) UserActivityFull(ctx context.Context, userID int64) (*model.UserActivityFull, error) {
	response, err := c.api.UserActivityFull(ctx, &clients.UserActivityFullRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc.Competition.UserActivityFull failed: %w", err)
	}

	return c.userActivityFullResponseToDto(response), nil
}

func (c *CompetitionClient) SaveSolution(ctx context.Context, userID, competitionID int64) error {
	_, err := c.api.SaveSolution(ctx, &clients.SaveSolutionRequest{
		UserID:        userID,
		CompetitionID: competitionID,
	})
	if err != nil {
		return fmt.Errorf("grpc.Competition.SaveSolution failed: %w", err)
	}

	return nil
}
