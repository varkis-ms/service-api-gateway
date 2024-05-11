package competitionrpc

import (
	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/pb/clients"
)

func (c *CompetitionClient) competitionListResponseToDto(competitionListResponse *clients.CompetitionListResponse) []model.CompetitionInfoSmall {
	compList := make([]model.CompetitionInfoSmall, len(competitionListResponse.CompetitionList))
	for idx, compProto := range competitionListResponse.CompetitionList {
		compList[idx] = model.CompetitionInfoSmall{
			CompetitionID: compProto.CompetitionID,
			Title:         compProto.Title,
			DatasetTitle:  compProto.DatasetTitle,
		}
	}

	return compList
}

func (c *CompetitionClient) competitionInfoResponseToDto(competitionInfoResponse *clients.CompetitionInfoResponse) *model.CompetitionInfo {
	return &model.CompetitionInfo{
		CompetitionID:      competitionInfoResponse.CompetitionID,
		Title:              competitionInfoResponse.Title,
		Description:        competitionInfoResponse.Description,
		DatasetTitle:       competitionInfoResponse.DatasetTitle,
		DatasetDescription: competitionInfoResponse.DatasetDescription,
		AmountUsers:        competitionInfoResponse.AmountUsers,
		MaximumScore:       competitionInfoResponse.MaximumScore,
	}
}

func (c *CompetitionClient) leaderboardResponseToDto(leaderboardResponse *clients.LeaderBoardResponse) *model.LeaderboardList {
	leaderboardList := make([]model.Leaderboard, len(leaderboardResponse.LeaderBoardList))
	for idx, leaderboardProto := range leaderboardResponse.LeaderBoardList {
		leaderboardList[idx] = model.Leaderboard{
			UserID:  leaderboardProto.UserID,
			Score:   leaderboardProto.Score,
			AddedAt: leaderboardProto.AddedAt,
		}
	}

	return &model.LeaderboardList{
		Leaderboard:   leaderboardList,
		CompetitionID: leaderboardResponse.CompetitionID,
	}
}

func (c *CompetitionClient) userActivityFullResponseToDto(userActivityFullResponse *clients.UserActivityFullResponse) *model.UserActivityFull {
	owner := make([]model.CompetitionInfoFullOwner, len(userActivityFullResponse.Owner))
	for idx, ownerProto := range userActivityFullResponse.Owner {
		owner[idx] = model.CompetitionInfoFullOwner{
			CompetitionID: ownerProto.CompetitionID,
			Title:         ownerProto.Title,
			DatasetTitle:  ownerProto.DatasetTitle,
			AmountUsers:   ownerProto.AmountUsers,
			AddedAt:       ownerProto.AddedAt,
		}
	}

	member := make([]model.CompetitionInfoFull, len(userActivityFullResponse.Member))
	for idx, memberProto := range userActivityFullResponse.Member {
		member[idx] = model.CompetitionInfoFull{
			CompetitionID: memberProto.CompetitionID,
			Title:         memberProto.Title,
			DatasetTitle:  memberProto.DatasetTitle,
			Score:         memberProto.Score,
			AddedAt:       memberProto.AddedAt,
		}
	}

	return &model.UserActivityFull{
		Owner:  owner,
		Member: member,
	}
}

func (c *CompetitionClient) userActivityTotalResponseToDto(userActivityTotalResponse *clients.UserActivityTotalResponse) *model.UserActivityTotal {
	return &model.UserActivityTotal{
		TotalTime:              userActivityTotalResponse.TotalTime,
		TotalAttempts:          userActivityTotalResponse.TotalAttempts,
		TotalCompetitions:      userActivityTotalResponse.TotalCompetitions,
		TotalOwnerCompetitions: userActivityTotalResponse.TotalOwnerCompetitions,
	}
}
