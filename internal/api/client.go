package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/andybalholm/brotli"
)

type ClaimResponse struct {
    Card struct {
        FutwizID     int    `json:"futwiz_id"`
        FullName     string `json:"full_name"`
        CardName     string `json:"card_name"`
        Nationality  string `json:"nationality"`
        TeamName     string `json:"team_name"`
        Rating       int    `json:"rating"`
        CardType     string `json:"card_type"`
        Position     string `json:"position"`
        League       string `json:"league"`
        Pace         int    `json:"pace"`
        Shooting     int    `json:"shooting"`
        Passing      int    `json:"passing"`
        Dribbling    int    `json:"dribbling"`
        Defending    int    `json:"defending"`
        Physicality  int    `json:"physicality"`
        Value        int    `json:"value"`
        Bin          int    `json:"bin"`
    } `json:"card"`
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func UnmarshalClaim(data []byte) (*ClaimResponse, error) {
	var result ClaimResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func SendRequest(reqDef apiRequest) (*ClaimResponse, error) {
	req, err := http.NewRequest(reqDef.Method, reqDef.URL, bytes.NewReader(reqDef.Body))
	if err != nil {
		return nil, err
	}

	for k, v := range reqDef.Headers {
		req.Header.Set(k, v)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var reader io.ReadCloser
    switch res.Header.Get("Content-Encoding") {
    case "gzip":
        reader, err = gzip.NewReader(res.Body)
        if err != nil {
            return nil, err
        }
        defer reader.Close()
    case "br":
        reader = io.NopCloser(brotli.NewReader(res.Body))
    default:
        reader = res.Body
    }

	var result ClaimResponse
	if err := json.NewDecoder(reader).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
