package main

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/chi07/go-comment-service/internal/http/handlers"
	"github.com/chi07/go-comment-service/internal/services"

	"github.com/chi07/go-comment-service/db"

	"github.com/chi07/go-comment-service/internal/repositories"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	viper.SetEnvPrefix("COMMENT")
	viper.AutomaticEnv()
	mongoURI := viper.GetString("MONGO_URI")
	dbName := viper.GetString("MONGO_DB_NAME")

	log.Info().Str("MONGO_URI", mongoURI).Msg("DB_URL")

	logLevelStr := viper.GetString("LOG_LEVEL")
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid LOG_LEVEL")
	}
	zerolog.LevelFieldName = "severity"
	zerolog.SetGlobalLevel(logLevel)
	log.Logger = log.With().Caller().Logger()

	mongoInstance, err := db.Connect(mongoURI, dbName)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to server")
	}

	defer func() {
		if err = mongoInstance.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	database := mongoInstance.Db

	// Create a Fiber app
	app := fiber.New()

	commentRepo := repositories.NewComment(database)
	createCommentService := services.NewCreateCommentService(commentRepo)
	updateCommentService := services.NewUpdateCommentService(commentRepo)
	deleteCommentService := services.NewDeleteCommentService(commentRepo)
	getCommentService := services.NewGetCommentService(commentRepo)

	voteRepo := repositories.NewVote(database)
	voteService := services.NewVoteCommentService(voteRepo)

	app.Post("/comments", handlers.NewCreateCommentHandler(createCommentService).CreateComment())
	app.Put("/comments/:id", handlers.NewUpdateCommentHandler(updateCommentService).UpdateComment())
	app.Delete("/comments/:id", handlers.NewDeleteCommentHandler(deleteCommentService).DeleteComment())
	app.Get("/comments/:id", handlers.NewGetCommentHandler(getCommentService).GetComment())

	app.Post("/votes", handlers.NewVoteHandler(voteService).Vote())

	err = app.Listen(":3000")
	log.Error().Err(err).Msg("Stopped serving HTTP")
}
