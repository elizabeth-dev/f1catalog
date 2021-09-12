package adapter

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/elizabeth-dev/f1catalog/api/app/query"
	"github.com/elizabeth-dev/f1catalog/api/domain/playlist"
	"github.com/pkg/errors"
)

type PlaylistModel struct {
	Data string `json:"data"`
}

type PlaylistRepository struct {
	dbClient *db.Client
}

func NewPlaylistRepository(dbClient *db.Client) PlaylistRepository {
	if dbClient == nil {
		panic("[PlaylistRepository] missing dbClient")
	}

	return PlaylistRepository{dbClient}
}

func (r PlaylistRepository) playlistRef() *db.Ref {
	return r.dbClient.NewRef("playlist")
}

func (r PlaylistRepository) GetPlaylist(ctx context.Context, contentId string) (*query.Playlist, error) {
	var playlistModel PlaylistModel

	if err := r.playlistRef().Child(contentId).Get(ctx, &playlistModel); err != nil {
		return nil, errors.Wrap(err, "[PlaylistRepository] Error retrieving playlist "+contentId)
	}

	pl, err := playlist.UnmarshalPlaylistFromDatabase(playlistModel.Data)

	if err != nil {
		return nil, err
	}

	return &query.Playlist{Data: pl.Data()}, nil
}

func (r PlaylistRepository) SetPlaylist(ctx context.Context, contentId string, pl *playlist.Playlist) error {
	playlistModel := r.marshalPlaylist(pl)

	return r.playlistRef().Child(contentId).Set(ctx, playlistModel)
}

func (r PlaylistRepository) marshalPlaylist(pl *playlist.Playlist) PlaylistModel {
	return PlaylistModel{Data: pl.Data()}
}
