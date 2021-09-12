package query

import (
	"context"
)

type GetPlaylistHandler struct {
	readModel GetPlaylistReadModel
}

type GetPlaylistReadModel interface {
	GetPlaylist(ctx context.Context, contentId string) (*Playlist, error)
}

func NewGetPlaylistHandler(readModel GetPlaylistReadModel) GetPlaylistHandler {
	if readModel == nil {
		panic("nil readModel")
	}
	return GetPlaylistHandler{readModel}
}

func (h GetPlaylistHandler) Handle(ctx context.Context, contentId string) (pl *Playlist, err error) {
	return h.readModel.GetPlaylist(ctx, contentId)
}
