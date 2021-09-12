package app

import (
	"github.com/elizabeth-dev/f1catalog/api/app/command"
	"github.com/elizabeth-dev/f1catalog/api/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SetPlaylist command.SetPlaylistHandler
}

type Queries struct {
	GetPlaylist    query.GetPlaylistHandler
	GetPlaylistURL query.GetPlaylistURLHandler
}
