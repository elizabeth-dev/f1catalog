package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	httpStatus int
}

func (e errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}

func HttpRespondWithError(err error, w http.ResponseWriter, r *http.Request, status int) {
	resp := errorResponse{status}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}
