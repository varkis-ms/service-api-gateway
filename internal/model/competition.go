package model

type CompetitionID struct {
	CompetitionID int64 `uri:"competitionID" binding:"required"`
}

type CompetitionCreate struct {
	UserID             int64
	Title              string `json:"title"               binding:"required"  example:"Competition Title"`
	Description        string `json:"description"         binding:"required"  example:"Competition Description"`
	DatasetTitle       string `json:"dataset_title"       binding:"required"  example:"Competition Description"`
	DatasetDescription string `json:"dataset_description" binding:"required"  example:"Competition Description"`
}

type CompetitionEdit struct {
	UserID             int64
	CompetitionID      int64
	Title              *string `json:"title"               binding:"omitempty"  example:"Competition Title"`
	Description        *string `json:"description"         binding:"omitempty"  example:"Competition Description"`
	DatasetTitle       *string `json:"dataset_title"       binding:"omitempty"  example:"Competition Description"`
	DatasetDescription *string `json:"dataset_description" binding:"omitempty"  example:"Competition Description"`
}

type CompetitionInfo struct {
	UserID             int64
	CompetitionID      int64
	Title              string
	Description        string
	DatasetTitle       string
	DatasetDescription string
	AmountUsers        int64
	MaximumScore       float32
}

type CompetitionInfoSmall struct {
	CompetitionID int64
	Title         string
	DatasetTitle  string
}

type LeaderboardList struct {
	Leaderboard   []Leaderboard
	CompetitionID int64
}

type Leaderboard struct {
	UserID  int64
	Score   float32
	AddedAt string
}

type FileUpload struct {
	UserID        int64
	CompetitionID int64
	FileType      int64
}

const (
	TypeNoData = iota
	TypeSolution
	TypeRequirements
	TypeDatasetTest
	TypeDatasetTrain
)
