package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/andybalholm/brotli"
	fhttp "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

type DailyResponse struct {
	Streak       int  `json:"streak"`
	IsPackReward bool `json:"is_pack_reward"`
}

type ClaimResponse struct {
	Card struct {
		FutwizID    int    `json:"futwiz_id"`
		FullName    string `json:"full_name"`
		CardName    string `json:"card_name"`
		Nationality string `json:"nationality"`
		TeamName    string `json:"team_name"`
		Rating      int    `json:"rating"`
		CardType    string `json:"card_type"`
		Position    string `json:"position"`
		League      string `json:"league"`
		Pace        int    `json:"pace"`
		Shooting    int    `json:"shooting"`
		Passing     int    `json:"passing"`
		Dribbling   int    `json:"dribbling"`
		Defending   int    `json:"defending"`
		Physicality int    `json:"physicality"`
		Value       int    `json:"value"`
		Bin         int    `json:"bin"`
		Uuid        string `json:"uuid"`
	} `json:"card"`
}

var httpClient tls_client.HttpClient

func init() {
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(20000),
		tls_client.WithClientProfile(profiles.Firefox_120), // mimick firefox fingerpring
		tls_client.WithNotFollowRedirects(),
	}
	client, _ := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	httpClient = client
}

func UnmarshalClaim(data []byte) (*ClaimResponse, error) {
	var result ClaimResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeBody(res *fhttp.Response) ([]byte, error) {
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	decoded := bodyBytes
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		gz, err := gzip.NewReader(bytes.NewReader(bodyBytes))
		if err == nil {
			defer gz.Close()
			tmp, err := io.ReadAll(gz)
			if err == nil {
				decoded = tmp
			}
		}
	case "br":
		br := brotli.NewReader(bytes.NewReader(bodyBytes))
		tmp, err := io.ReadAll(br)
		if err == nil {
			decoded = tmp
		}
	}

	return decoded, nil
}

func SendRequest(reqDef apiRequest) (*ClaimResponse, error) {
	req, err := fhttp.NewRequest(reqDef.Method, reqDef.URL, bytes.NewReader(reqDef.Body))
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

	decoded, err := decodeBody(res)
	if err != nil {
		return nil, err
	}

	var result ClaimResponse
	if err := json.Unmarshal(decoded, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func DailyRequest(reqDef dailyRequest) (*DailyResponse, error) {
	req, err := fhttp.NewRequest(reqDef.Method, reqDef.URL, bytes.NewReader(reqDef.Body))
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

	decoded, err := decodeBody(res)
	if err != nil {
		return nil, err
	}

	var result DailyResponse
	if err := json.Unmarshal(decoded, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
