package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type ClaimResponse struct {
	Card struct {
		FullName string `json:"full_name"`
		Position string `json:"position"`
		Rating   int    `json:"rating"`
		Value    int    `json:"value"`
		Bin      int    `json:"bin"`
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
	default:
		reader = res.Body
	}

	var result ClaimResponse
	if err := json.NewDecoder(reader).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
