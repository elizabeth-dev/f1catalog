package main

import (
	"io"
	"log"
	"net/http"
)

func getPlayToken(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == "playToken" {
			return cookie
		}
	}

	return nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", relay)

	log.Printf("Running...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func relay(w http.ResponseWriter, r *http.Request) {
	var targetHost string
	query := r.URL.Query()
	targetHost = query.Get("host")
	query.Del("host")
	r.URL.RawQuery = query.Encode()

	hostCookie, err := r.Cookie("host")

	if err == nil && hostCookie != nil {
		targetHost = hostCookie.Value
	}

	r.Host = targetHost
	r.URL.Scheme = "https"
	r.URL.Host = targetHost
	r.RequestURI = ""

	res, err := http.DefaultClient.Do(r)

	if err != nil {
		log.Printf("Error while calling OTT")
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playTokenCookie := getPlayToken(res.Cookies())
	if playTokenCookie != nil {
		playTokenCookie.Secure = false
		http.SetCookie(w, playTokenCookie)

		if hostCookie == nil {
			http.SetCookie(w, &http.Cookie{Name: "host", Value: targetHost, Path: playTokenCookie.Path, SameSite: playTokenCookie.SameSite})
		}
	}

	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", res.Header.Get("Content-Length"))
	w.Header().Set("Content-Range", res.Header.Get("Content-Range"))
	w.WriteHeader(res.StatusCode)

	_, err = io.Copy(w, res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = res.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
