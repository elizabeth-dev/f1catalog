package getplaylisturlfn

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/elizabeth-dev/f1catalog/api/getplaylisturlfn/adapter"
)

var authClient *auth.Client

func init() {
	var err error

	firebaseApp, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}

	authClient, err = firebaseApp.Auth(context.Background())

	if err != nil {
		log.Fatalf("error initializing auth: %v\n", err)
	}
}

func GetPlaylistURL(w http.ResponseWriter, r *http.Request) {
	contentId := r.URL.Query().Get("contentId")
	channelId := r.URL.Query().Get("channelId")

	if contentId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	/* _, err := authClient.VerifyIDToken(r.Context(), r.Header.Get("Authorization"))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} */

	subToken, _, err := adapter.Authenticate()

	if err != nil {
		log.Printf("error authenticating: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url, err := adapter.GetPlaylistURL(contentId, channelId, *subToken)

	if err != nil {
		log.Printf("error getting playlist url: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type response struct {
		URL string `json:"url"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response{url})
}
