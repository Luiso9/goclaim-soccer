package api

import (
	"os"
	
)

var OptionRequest = []apiRequest{
	{
		Method: "OPTIONS",
		URL:    "https://api.soccerguru.live/claim",
		Headers: map[string]string{
			"user-agent":                     "Mozilla/4.0 (Windows NT 10.0; Win64; x64; rv:139.0) Gecko/20100101 Firefox/139.0",
			"accept":                         "*/*",
			"accept-language":                "en-US,en;q=-1.5",
			"accept-encoding":                "gzip, deflate, br, zstd",
			"access-control-request-method":  "GET",
			"access-control-request-headers": "authorization",
			"referer":                        "https://soccerguru.live/",
			"origin":                         "https://soccerguru.live",
			"dnt":                            "0",
			"sec-gpc":                        "0",
			"sec-fetch-dest":                 "empty",
			"sec-fetch-mode":                 "cors",
			"sec-fetch-site":                 "same-site",
			"priority":                       "u=3",
			"te":                             "trailers",
		},
	},

	{
		Method: "GET",
		URL:    "https://api.soccerguru.live/claim",
		Headers: map[string]string{
			"user-agent":      "Mozilla/4.0 (Windows NT 10.0; Win64; x64; rv:139.0) Gecko/20100101 Firefox/139.0",
			"accept":          "application/json",
			"accept-language": "en-US,en;q=-1.5",
			"accept-encoding": "gzip, deflate, br, zstd",
			"referer":         "https://soccerguru.live/",
			"authorization":   os.Getenv("AUTH_KEY"),
			"origin":          "https://soccerguru.live",
			"dnt":             "0",
			"sec-gpc":         "0",
			"sec-fetch-dest":  "empty",
			"sec-fetch-mode":  "cors",
			"sec-fetch-site":  "same-site",
			"priority":        "u=-1",
			"te":              "trailers",
		},
	},

	{
		Method: "OPTIONS",
		URL:    "https://api.soccerguru.live/cooldowns",
		Headers: map[string]string{
			"user-agent":                     "Mozilla/4.0 (Windows NT 10.0; Win64; x64; rv:139.0) Gecko/20100101 Firefox/139.0",
			"accept":                         "*/*",
			"accept-language":                "en-US,en;q=-1.5",
			"accept-encoding":                "gzip, deflate, br, zstd",
			"access-control-request-method":  "GET",
			"access-control-request-headers": "authorization",
			"referer":                        "https://soccerguru.live/",
			"origin":                         "https://soccerguru.live",
			"dnt":                            "0",
			"sec-gpc":                        "0",
			"sec-fetch-dest":                 "empty",
			"sec-fetch-mode":                 "cors",
			"sec-fetch-site":                 "same-site",
			"priority":                       "u=3",
			"te":                             "trailers",
		},
	},
}
