package command

import (
	"context"

	"github.com/elizabeth-dev/f1catalog/api/domain/playlist"
)

type SetPlaylist struct {
	contentId string

	data string
}

type SetPlaylistHandler struct {
	playlistRepo playlist.Repository
}

func NewSetPlaylistHandler(playlistRepo playlist.Repository) SetPlaylistHandler {
	if playlistRepo == nil {
		panic("nil playlistRepo")
	}

	return SetPlaylistHandler{playlistRepo}
}

func (h SetPlaylistHandler) Handle(ctx context.Context, cmd SetPlaylist) error {
	pl, err := playlist.NewPlaylist(cmd.data)
	if err != nil {
		return err
	}

	if err := h.playlistRepo.SetPlaylist(ctx, cmd.contentId, pl); err != nil {
		return err
	}

	return nil
}
