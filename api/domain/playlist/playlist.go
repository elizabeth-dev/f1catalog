package playlist

import "github.com/pkg/errors"

type Playlist struct {
	data string
}

func (p *Playlist) Data() string {
	return p.data
}

type Factory struct {
}

func NewFactory() (Factory, error) {
	return Factory{}, nil
}

func (f Factory) IsZero() bool {
	return f == Factory{}
}

func NewPlaylist(data string) (*Playlist, error) {
	if data == "" {
		return nil, errors.New("[Playlist] Empty data")
	}

	return &Playlist{data}, nil
}

func UnmarshalPlaylistFromDatabase(data string) (*Playlist, error) {
	return &Playlist{data}, nil
}
