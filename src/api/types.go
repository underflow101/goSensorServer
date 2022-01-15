package api

import (
	"time"

	"github.com/underflow101/src/db"
)

type ApiResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type ApiGpsResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Body       db.Gps `json:"body"`
}

func GetTimeNowUnix() int64 {
	now := time.Now()
	unixNow := now.Unix()

	return unixNow
}
