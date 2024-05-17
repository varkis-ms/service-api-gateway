package app

import (
	"log/slog"
	"os"

	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_create"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_edit"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/competition_list_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/download_file"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/leaderboard_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/login"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/profile_edit"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/profile_get"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/save_solution"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/signup"
	"github.com/varkis-ms/service-api-gateway/internal/api/handler/upload_file"
	mwauth "github.com/varkis-ms/service-api-gateway/internal/api/middleware/auth"
	httpapp "github.com/varkis-ms/service-api-gateway/internal/app/http"
	"github.com/varkis-ms/service-api-gateway/internal/clients/rpc"
	authclient "github.com/varkis-ms/service-api-gateway/internal/clients/rpc/auth"
	competitionclient "github.com/varkis-ms/service-api-gateway/internal/clients/rpc/competition"
	solutionclient "github.com/varkis-ms/service-api-gateway/internal/clients/rpc/solution"
	userinfoclient "github.com/varkis-ms/service-api-gateway/internal/clients/rpc/user_info"
	"github.com/varkis-ms/service-api-gateway/internal/config"
	"github.com/varkis-ms/service-api-gateway/internal/pkg/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run(configPath string) {
	// Config
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic("config.LoadConfig failed" + err.Error())
	}

	logger := setupLogger(cfg.Env)

	// Clients
	authClient := authclient.New(rpc.RegistrationClient(cfg.AuthClientAddr, logger))
	userInfoClient := userinfoclient.New(rpc.RegistrationClient(cfg.UserInfoClientAddr, logger))
	competitionClient := competitionclient.New(rpc.RegistrationClient(cfg.CompetitionClientAddr, logger))
	solutionClient := solutionclient.New(rpc.RegistrationClient(cfg.SolutionClientAddr, logger))

	// Handler
	logger.Info("Initializing handlers and routes...")
	signupHandler := signup.New(authClient, userInfoClient, logger)
	loginHandler := login.New(authClient, logger)

	profileEditHandler := profile_edit.New(userInfoClient, logger)
	profileGetHandler := profile_get.New(userInfoClient, competitionClient, logger)

	competitionCreateHandler := competition_create.New(competitionClient, logger)
	competitionEditHandler := competition_edit.New(competitionClient, logger)
	competitionGetHandler := competition_get.New(competitionClient, logger)
	competitionListGetHandler := competition_list_get.New(competitionClient, logger)
	leaderboardGetHandler := leaderboard_get.New(competitionClient, logger)
	uploadFileHandler := upload_file.New(solutionClient, logger)
	downloadFileHandler := download_file.New(solutionClient, logger)
	saveSolutionHandler := save_solution.New(competitionClient, logger)

	// Middlewares
	mw := mwauth.New(authClient, logger)

	// Registration
	httpapp.NewRouter(
		cfg.PortHttp,
		mw,
		signupHandler,
		loginHandler,
		profileEditHandler,
		profileGetHandler,
		competitionCreateHandler,
		competitionEditHandler,
		competitionGetHandler,
		competitionListGetHandler,
		leaderboardGetHandler,
		uploadFileHandler,
		downloadFileHandler,
		saveSolutionHandler,
		logger,
	)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
