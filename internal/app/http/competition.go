package http

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_create"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_edit"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_list_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/download_file"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/leaderboard_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/upload_file"

	mwauth "github.com/varkis-ms/service-api-gateway/internal/api/middleware/auth"
	"github.com/varkis-ms/service-api-gateway/internal/model"
)

type competitionRoutes struct {
	competitionCreateHandler  *competition_create.Handler
	competitionEditHandler    *competition_edit.Handler
	competitionGetHandler     *competition_get.Handler
	competitionListGetHandler *competition_list_get.Handler
	leaderboardGetHandler     *leaderboard_get.Handler
	uploadFileHandler         *upload_file.Handler
	downloadFileHandler       *download_file.Handler
	log                       *slog.Logger
}

func newCompetitionRoutes(
	h *gin.RouterGroup,
	mw *mwauth.Middleware,
	competitionCreateHandler *competition_create.Handler,
	competitionEditHandler *competition_edit.Handler,
	competitionGetHandler *competition_get.Handler,
	competitionListGetHandler *competition_list_get.Handler,
	leaderboardGetHandler *leaderboard_get.Handler,
	uploadFileHandler *upload_file.Handler,
	downloadFileHandler *download_file.Handler,
	log *slog.Logger,
) {
	r := &competitionRoutes{
		competitionCreateHandler:  competitionCreateHandler,
		competitionEditHandler:    competitionEditHandler,
		competitionGetHandler:     competitionGetHandler,
		competitionListGetHandler: competitionListGetHandler,
		leaderboardGetHandler:     leaderboardGetHandler,
		uploadFileHandler:         uploadFileHandler,
		downloadFileHandler:       downloadFileHandler,
		log:                       log,
	}

	{
		h.POST("/create", mw.Middleware, r.competitionCreate)
		h.POST("/:competitionID/edit", mw.Middleware, r.competitionEdit)
		h.GET("/list", mw.Middleware, r.competitionList)
		h.GET("/:competitionID", mw.Middleware, r.competitionInfo)
		h.GET("/:competitionID/leaderboard", r.leaderboard)
		h.POST("/:competitionID/upload", mw.Middleware, r.upload)
		h.GET("/:competitionID/download", r.download)
	}
}

func (r *competitionRoutes) competitionCreate(c *gin.Context) {
	var req *model.CompetitionCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)
		return
	}

	req.UserID = c.GetInt64("userID")
	competitionID, err := r.competitionCreateHandler.Handle(c, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"competitionID": competitionID,
	})
}

func (r *competitionRoutes) competitionEdit(c *gin.Context) {
	var compID model.CompetitionID
	if err := c.ShouldBindUri(&compID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)
		return
	}

	var req *model.CompetitionEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)
		return
	}

	req.UserID = c.GetInt64("userID")
	req.CompetitionID = compID.CompetitionID
	err := r.competitionEditHandler.Handle(c, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (r *competitionRoutes) competitionInfo(c *gin.Context) {
	var compID model.CompetitionID
	if err := c.ShouldBindUri(&compID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)

		return
	}

	compList, err := r.competitionGetHandler.Handle(c, c.GetInt64("userID"), compID.CompetitionID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, compList)
}

func (r *competitionRoutes) competitionList(c *gin.Context) {
	compList, err := r.competitionListGetHandler.Handle(c, c.GetInt64("userID"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, compList)
}

func (r *competitionRoutes) leaderboard(c *gin.Context) {
	var compID model.CompetitionID
	if err := c.ShouldBindUri(&compID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)

		return
	}
	leaderboards, err := r.leaderboardGetHandler.Handle(c, compID.CompetitionID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, leaderboards)
}

func (r *competitionRoutes) upload(c *gin.Context) {
	var compID model.CompetitionID
	if err := c.ShouldBindUri(&compID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)

		return
	}

	userID := c.GetInt64("userID")
	_, ok := c.GetPostForm("solve")
	if ok {
		// TODO: ENUM
		scriptFile, err := c.FormFile("script")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)

			return
		}

		if err = r.uploadFileHandler.Handle(c, scriptFile, userID, compID.CompetitionID, model.TypeSolution); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		reqFile, err := c.FormFile("req")
		if err == nil {
			if err = r.uploadFileHandler.Handle(c, reqFile, userID, compID.CompetitionID, model.TypeRequirements); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)

				return
			}
		}
	} else {
		datasetTestFile, err := c.FormFile("dataset_test")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)

			return
		}
		datasetTrainFile, err := c.FormFile("dataset_test")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)

			return
		}

		if err = r.uploadFileHandler.Handle(c, datasetTestFile, userID, compID.CompetitionID, model.TypeDatasetTest); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		if err = r.uploadFileHandler.Handle(c, datasetTrainFile, userID, compID.CompetitionID, model.TypeDatasetTrain); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (r *competitionRoutes) download(c *gin.Context) {
	var compID model.CompetitionID
	if err := c.ShouldBindUri(&compID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)
		return
	}

	dataset, err := r.downloadFileHandler.Handle(c, compID.CompetitionID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}
	defer dataset.Reset()

	c.Writer.Header().Set("Content-Type", "text/csv")
	// Set the content disposition header to force download
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=dataset.csv")
	io.Copy(c.Writer, dataset)
}
