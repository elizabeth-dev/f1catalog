package playlist

import "context"

type Repository interface {
	SetPlaylist(ctx context.Context, contentId string, playlist *Playlist) error
}
