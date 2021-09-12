package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/elizabeth-dev/f1catalog/api/adapter"
	"github.com/elizabeth-dev/f1catalog/api/app"
	"github.com/elizabeth-dev/f1catalog/api/app/command"
	"github.com/elizabeth-dev/f1catalog/api/app/query"
	"google.golang.org/api/option"
)

func NewApplication(ctx context.Context) app.Application {
	f1TVClient := adapter.NewF1TVClient()

	return newApplication(ctx, f1TVClient)
}

func newApplication(ctx context.Context, f1TVClient command.F1TVService) app.Application {
	firebaseApp, err := firebase.NewApp(context.Background(), &firebase.Config{DatabaseURL: "https://f1tv-f515e-default-rtdb.firebaseio.com"}, option.WithCredentialsFile("aaa.json"))
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		panic(err)
	}

	dbClient, err := firebaseApp.Database(ctx)
	if err != nil {
		log.Fatalf("Error initializing database client: %v\n", err)
		panic(err)
	}

	playlistRepository := adapter.NewPlaylistRepository(dbClient)

	return app.Application{
		Queries: app.Queries{
			GetPlaylist:    query.NewGetPlaylistHandler(playlistRepository),
			GetPlaylistURL: query.NewGetPlaylistURLHandler(f1TVClient),
		},
		Commands: app.Commands{
			SetPlaylist: command.NewSetPlaylistHandler(playlistRepository),
		},
	}
}
