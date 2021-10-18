package corsproxyfn

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/elizabeth-dev/f1catalog/api/corsproxyfn/adapter"
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

func CorsProxy(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS requests
	if r.Method == "OPTIONS" {
		for n, h := range r.Header {
			if strings.Contains(n, "Access-Control-Request") {
				for _, h := range h {
					k := strings.Replace(n, "Request", "Allow", 1)
					w.Header().Add(k, h)
				}
			}
		}
		return
	}

	uri := r.URL.Query().Get("uri")
	if uri == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check user auth
	/* _, err := authClient.VerifyIDToken(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} */

	resBody, err := adapter.ProxyReq(r.Method, uri, r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.Copy(w, resBody)
}
