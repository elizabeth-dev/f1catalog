package adapter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	baseURL = "https://f1tv.formula1.com"
	authURL = "https://api.formula1.com/v2/account/subscriber/authenticate/by-password"

	playbackRequestPath = "/1.0/R/ENG/BIG_SCREEN_HLS/ALL/CONTENT/PLAY?contentId="

	apiKey = "fCUCjWrKPu9ylJwRAv8BpGLEgiAuThx7"
)

type F1TVClient struct {
	HttpClient *http.Client
}

func NewF1TVClient() F1TVClient {
	return F1TVClient{HttpClient: http.DefaultClient}
}

func Authenticate() (*string, *int64, error) {
	type request struct {
		Login    string `json:"Login"`
		Password string `json:"Password"`
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(request{Login: "elizabethmc1999+f1tvfrance@gmail.com", Password: "Lr9XRS56kuSugAZ"})
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, authURL, payloadBuf)
	req.Header.Set("apiKey", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "RaceControl f1viewer")

	if err != nil {
		return nil, nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var auth struct {
		Data struct {
			SubscriptionStatus string `json:"subscriptionStatus"`
			SubscriptionToken  string `json:"subscriptionToken"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&auth)

	if err != nil {
		return nil, nil, errors.New("[F1TVClient] Error parsing response")
	}

	if auth.Data.SubscriptionToken == "" {
		return nil, nil, errors.New("[F1TVClient] Could not get subscription token")
	}

	jwtStr, err := base64.RawStdEncoding.DecodeString(strings.Split(auth.Data.SubscriptionToken, ".")[1])

	var jwt struct {
		Exp int64 `json:"exp"`
	}

	json.Unmarshal([]byte(jwtStr), &jwt)

	return &auth.Data.SubscriptionToken, &jwt.Exp, err
}

func GetPlaylistURL(contentId string, channelId string, subToken string) (string, error) {
	url := baseURL + playbackRequestPath + contentId

	if channelId != "" {
		url += "&channelId=" + channelId
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", nil
	}

	req.Header.Set("ascendontoken", subToken)
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var resp struct {
		ResultCode       string `json:"resultCode"`
		Message          string `json:"message"`
		ErrorDescription string `json:"errorDescription"`
		ResultObj        struct {
			EntitlementToken string `json:"entitlementToken"`
			URL              string `json:"url"`
			StreamType       string `json:"streamType"`
		} `json:"resultObj"`
		SystemTime int `json:"systemTime"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return "", err
	}

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		err = errors.New(resp.Message)
	} else if resp.ResultObj.URL == "" {
		err = fmt.Errorf("API returned empty URL: %s", resp.Message)
	}

	return resp.ResultObj.URL, err
}
