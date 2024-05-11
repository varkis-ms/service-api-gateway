package http

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_create"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_edit"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_list_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/download_file"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/leaderboard_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/login"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/profile_edit"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/profile_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/signup"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/upload_file"
	"github.com/varkis-ms/service-api-gateway/internal/api/middleware/auth"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/httpserver"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/sl"
)

func NewRouter(
	portHttp string,
	mw *mwauth.Middleware,
	signupHandler *signup.Handler,
	loginHandler *login.Handler,
	profileEditHandler *profile_edit.Handler,
	profileGetHandler *profile_get.Handler,
	competitionCreateHandler *competition_create.Handler,
	competitionEditHandler *competition_edit.Handler,
	competitionGetHandler *competition_get.Handler,
	competitionListGetHandler *competition_list_get.Handler,
	leaderboardGetHandler *leaderboard_get.Handler,
	uploadFileHandler *upload_file.Handler,
	downloadFileHandler *download_file.Handler,
	log *slog.Logger,
) {
	handler := gin.Default()
	// Swagger
	//docs.SwaggerInfo.BasePath = "/api/v1"
	//handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routers
	h := handler.Group("/api/v1")
	{
		newUserRoutes(
			h.Group("/user"),
			signupHandler,
			loginHandler,
			log,
		)
		newProfileRoutes(
			h.Group("/profile"),
			mw,
			profileEditHandler,
			profileGetHandler,
			log,
		)
		newCompetitionRoutes(
			h.Group("/competition"),
			mw,
			competitionCreateHandler,
			competitionEditHandler,
			competitionGetHandler,
			competitionListGetHandler,
			leaderboardGetHandler,
			uploadFileHandler,
			downloadFileHandler,
			log,
		)
	}

	// HTTP server
	log.Info("Starting http server on port :%s", slog.String("addr", portHttp))
	httpServer := httpserver.New(handler, httpserver.Port(fmt.Sprintf(portHttp)))

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app.Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Error("app.Run - httpServer.Notify", sl.Err(err))
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err := httpServer.Shutdown()
	if err != nil {
		log.Error("app.Run - httpServer.Shutdown", sl.Err(err))
	}
}
