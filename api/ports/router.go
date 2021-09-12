package ports

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ServerInterface interface {
	//  (GET //playlist/{contentId}/url)
	GetPlaylistURL(w http.ResponseWriter, r *http.Request)
}

func GetPlaylistURLCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	r.Group(func(r chi.Router) {
		r.Use(GetPlaylistURLCtx)
		r.Get("/playlist/{contentId}/url", si.GetPlaylistURL)
	})

	return r
}
