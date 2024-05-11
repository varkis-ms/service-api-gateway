package userinforpc

import (
	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/pb/clients"
)

func (c *UserInfoClient) userInfoGetResponseToDto(userInfoGetResponse *clients.UserInfoGetResponse) *model.UserInfo {
	return &model.UserInfo{
		UserID:      userInfoGetResponse.UserID,
		Nickname:    userInfoGetResponse.Nickname,
		FullName:    userInfoGetResponse.FullName,
		Birthday:    userInfoGetResponse.Birthday,
		Location:    userInfoGetResponse.Location,
		Description: userInfoGetResponse.Description,
	}
}
