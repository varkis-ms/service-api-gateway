package userinforpc

import (
	"context"
	"fmt"

	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/pb/clients"
	"google.golang.org/grpc"
)

type UserInfoClient struct {
	api clients.UserInfoClient
}

func New(
	rpcClient *grpc.ClientConn,
) *UserInfoClient {
	return &UserInfoClient{
		api: clients.NewUserInfoClient(rpcClient),
	}
}

func (c *UserInfoClient) UserInfoSave(ctx context.Context, userID int64, email string) error {
	_, err := c.api.UserInfoSave(ctx, &clients.UserInfoSaveRequest{
		UserID: userID,
		Email:  email,
	})
	if err != nil {
		return fmt.Errorf("grpc.UserInfo.UserInfoSave failed: %w", err)
	}

	return nil
}

func (c *UserInfoClient) UserInfoEdit(ctx context.Context, userInfo *model.UserInfo) error {
	_, err := c.api.UserInfoEdit(ctx, &clients.UserInfoEditRequest{
		UserID:      userInfo.UserID,
		Nickname:    userInfo.Nickname,
		FullName:    userInfo.FullName,
		Birthday:    userInfo.Birthday,
		Location:    userInfo.Location,
		Description: userInfo.Description,
	})
	if err != nil {
		return fmt.Errorf("grpc.UserInfo.UserInfoSave failed: %w", err)
	}

	return nil
}

func (c *UserInfoClient) UserInfoGet(ctx context.Context, userID int64) (*model.UserInfo, error) {
	userInfoGetResponse, err := c.api.UserInfoGet(ctx, &clients.UserInfoGetRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc.UserInfo.UserInfoSave failed: %w", err)
	}

	return c.userInfoGetResponseToDto(userInfoGetResponse), nil
}
