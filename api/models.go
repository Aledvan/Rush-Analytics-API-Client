package api

import _ "time"

type RankTrackerData struct {
	Date  string `json:"date"`
	Ranks []struct {
		Keyword string `json:"keyword"`
		Rank    int    `json:"rank"`
		Url     string `json:"url"`
	} `json:"ranks"`
}

type RankTrackerResponse struct {
	Success bool              `json:"success"`
	Data    []RankTrackerData `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
	Error   string            `json:"error,omitempty"` // Некоторые API возвращают error вместо message
}

type RawResponse struct {
	Body string
}
