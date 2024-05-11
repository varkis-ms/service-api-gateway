package authrpc

import (
	"context"
	"fmt"

	"github.com/varkis-ms/service-api-gateway/internal/pkg/pb/clients"
	"google.golang.org/grpc"
)

type AuthClient struct {
	api clients.AuthClient
}

func New(
	rpcClient *grpc.ClientConn,
) *AuthClient {
	return &AuthClient{
		api: clients.NewAuthClient(rpcClient),
	}
}

func (c *AuthClient) Signup(ctx context.Context, email, password string) (int64, error) {
	response, err := c.api.Signup(ctx, &clients.SignupRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return 0, fmt.Errorf("grpc.Auth.Signup failed: %w", err)
	}

	return response.UserID, nil
}

func (c *AuthClient) Login(ctx context.Context, email, password string) (string, error) {
	response, err := c.api.Login(ctx, &clients.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", fmt.Errorf("grpc.Auth.Login failed: %w", err)
	}

	return response.Token, nil
}

func (c *AuthClient) Validate(ctx context.Context, token string) (int64, error) {
	response, err := c.api.Validate(ctx, &clients.ValidateRequest{
		Token: token,
	})
	if err != nil {
		return 0, fmt.Errorf("grpc.Auth.Validate failed: %w", err)
	}

	return response.UserID, nil
}
