package getplaylisturlfn

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/elizabeth-dev/f1catalog/api/fetchplaybackfn/adapter"
)

var authClient *auth.Client
var token *string
var tokenExp *int64

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

func getSubscriptionToken() (*string, error) {
	if tokenExp == nil || time.Now().Unix() > *tokenExp {
		_token, _tokenExp, err := adapter.Authenticate()
		if err != nil {
			return nil, err
		}

		token, tokenExp = _token, _tokenExp
	}

	return token, nil
}

func FetchPlayback(w http.ResponseWriter, r *http.Request) {
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

	/* Authenticating to F1TV (if necessary) */
	subToken, err := getSubscriptionToken()

	if err != nil {
		log.Printf("error authenticating: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/* Retrieving CDN URL */
	playbackUrl, err := adapter.GetPlaylistURL(contentId, channelId, *subToken)

	if err != nil {
		log.Printf("error getting playlist url: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/* Removing start parameter */
	parsedUrl, err := url.Parse(playbackUrl)

	if err != nil {
		log.Printf("error parsing playlist url: %v\n", err)
		http.Redirect(w, r, playbackUrl, http.StatusTemporaryRedirect)
		return
	}

	query := parsedUrl.Query()
	query.Del("start")
	parsedUrl.RawQuery = query.Encode()

	/* Sending HTTP 307 to CDN URL*/
	http.Redirect(w, r, parsedUrl.String(), http.StatusTemporaryRedirect)
}
