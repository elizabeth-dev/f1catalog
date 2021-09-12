package ports

import (
	"log"
	"net/http"

	"github.com/elizabeth-dev/f1catalog/api/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type HttpServer struct {
	application app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{application}
}

func (s HttpServer) GetPlaylistURL(w http.ResponseWriter, r *http.Request) {
	contentId := chi.URLParam(r, "contentId")

	url, err := s.application.Queries.GetPlaylistURL.Handle(r.Context(), contentId)
	if err != nil {
		log.Print(err.Error())
		HttpRespondWithError(err, w, r, 500)
		return
	}

	resp := PlaylistURL{url}
	render.Respond(w, r, resp)
}
