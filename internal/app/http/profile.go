package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/profile_edit"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/profile_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/middleware/auth"
	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type profileRoutes struct {
	profileEditHandler *profile_edit.Handler
	profileGetHandler  *profile_get.Handler
	log                *slog.Logger
}

func newProfileRoutes(
	h *gin.RouterGroup,
	mw *mwauth.Middleware,
	userInfoEditHandler *profile_edit.Handler,
	profileGetHandler *profile_get.Handler,
	log *slog.Logger,
) {
	r := &profileRoutes{
		profileEditHandler: userInfoEditHandler,
		profileGetHandler:  profileGetHandler,
		log:                log,
	}

	{
		h.POST("/edit", mw.Middleware, r.edit)
		h.GET("/me", mw.Middleware, r.profileMe)
		h.GET("/:userID", mw.Middleware, r.profile)
	}
}

func (r *profileRoutes) edit(c *gin.Context) {
	var req *model.UserInfo
	if err := c.Bind(&req); err != nil {
		r.log.Error("bad request", sl.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	req.UserID = c.GetInt64("userID")
	err := r.profileEditHandler.Handle(c, req)
	if err != nil {
		r.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, nil)
}

func (r *profileRoutes) profileMe(c *gin.Context) {
	userID := c.GetInt64("userID")
	profile, err := r.profileGetHandler.Handle(c, userID)
	if err != nil {
		r.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, profile)
}

func (r *profileRoutes) profile(c *gin.Context) {
	var profileGetRequest *model.ProfileGetRequest
	if err := c.ShouldBindUri(&profileGetRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrBadRequest)
		return
	}

	profile, err := r.profileGetHandler.Handle(c, profileGetRequest.UserID)
	if err != nil {
		r.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, profile)
}
