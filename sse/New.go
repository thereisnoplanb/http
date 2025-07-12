package sse

import (
	"net/http"

	"github.com/thereisnoplanb/http/statusRecorder"
)

func New(response http.ResponseWriter) SSE {
	response.Header().Set("Content-Type", "text/event-stream")
	response.Header().Set("Cache-Control", "no-cache")
	response.Header().Set("Connection", "keep-alive")

	return &sse{
		ResponseWriter: statusRecorder.New(response),
	}
}
