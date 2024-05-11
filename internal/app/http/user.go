package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/login"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/signup"
	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type userRoutes struct {
	signupHandler *signup.Handler
	loginHandler  *login.Handler
	log           *slog.Logger
}

func newUserRoutes(
	h *gin.RouterGroup,
	signupHandler *signup.Handler,
	loginHandler *login.Handler,
	log *slog.Logger,
) {
	r := &userRoutes{
		signupHandler: signupHandler,
		loginHandler:  loginHandler,
		log:           log,
	}

	{
		h.POST("/signup", r.signup)
		h.POST("/login", r.login)
	}
}

func (r *userRoutes) signup(c *gin.Context) {
	var req *model.SignUpRequest
	if err := c.Bind(&req); err != nil {
		r.log.Error("bad request", sl.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	err := r.signupHandler.Handle(c, req)
	if err != nil {
		r.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (r *userRoutes) login(c *gin.Context) {
	// TODO: change model
	var req *model.SignUpRequest
	if err := c.Bind(&req); err != nil {
		r.log.Error("bad request", sl.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	token, err := r.loginHandler.Handle(c, req)
	if err != nil {
		r.log.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
