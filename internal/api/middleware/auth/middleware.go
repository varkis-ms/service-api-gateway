package mwauth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/varkis-ms/service-api-gateway/internal/model"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

type Middleware struct {
	authClient AuthClient
	log        *slog.Logger
}

func New(
	authClient AuthClient,
	log *slog.Logger,
) *Middleware {
	return &Middleware{
		authClient: authClient,
		log:        log,
	}
}

func (m *Middleware) Middleware(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		m.log.Error("no cookie")
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrNoUser)

		return
	}

	userID, err := m.authClient.Validate(c, token)
	if err != nil {
		m.log.Error("authClient.Validate failed", sl.Err(err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrNoUser)

		return
	}

	c.Set("userID", userID)
	c.Next()
}
