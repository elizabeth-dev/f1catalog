package query

import (
	"context"
	"log"
	"time"
)

type GetPlaylistURLHandler struct {
	subToken  string
	subExp    uint
	readModel GetPlaylistURLReadModel
}

type GetPlaylistURLReadModel interface {
	GetPlaylistURL(contentId string, subToken string) (string, error)
	Authenticate() (*string, *uint, error)
}

func NewGetPlaylistURLHandler(readModel GetPlaylistURLReadModel) GetPlaylistURLHandler {
	if readModel == nil {
		panic("nil readModel")
	}
	return GetPlaylistURLHandler{readModel: readModel, subExp: 0, subToken: ""}
}

func (h GetPlaylistURLHandler) Handle(ctx context.Context, contentId string) (url string, err error) {
	if h.subExp <= uint(time.Now().Unix()) {
		subToken, subExp, err := h.readModel.Authenticate()

		if err != nil {
			log.Print(err.Error())

			panic("F1TV auth error")
		}

		h.subToken = *subToken
		h.subExp = *subExp
	}

	return h.readModel.GetPlaylistURL(contentId, h.subToken)
}
