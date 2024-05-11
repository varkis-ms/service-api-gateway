package model

type UserInfo struct {
	UserID      int64
	Nickname    *string
	FullName    *string
	Birthday    *string
	Location    *string
	Description *string
	CreatedAt   *string
}

type UserActivityTotal struct {
	TotalTime              string
	TotalAttempts          int64
	TotalCompetitions      int64
	TotalOwnerCompetitions int64
}

type CompetitionInfoFull struct {
	CompetitionID int64
	Title         string
	DatasetTitle  string
	Score         float32
	AddedAt       string
}

type CompetitionInfoFullOwner struct {
	CompetitionID int64
	Title         string
	DatasetTitle  string
	AmountUsers   int64
	AddedAt       string
}

type UserActivityFull struct {
	Owner  []CompetitionInfoFullOwner
	Member []CompetitionInfoFull
}

type SignUpRequest struct {
	Email    string
	Password string
}
