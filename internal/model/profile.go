package model

type ProfileGetRequest struct {
	UserID int64 `uri:"userID" binding:"required"`
}

type ProfileGetResponse struct {
	UserMainInfo          UserInfo
	UserActivityFullInfo  UserActivityFull
	UserActivityTotalInfo UserActivityTotal
}
