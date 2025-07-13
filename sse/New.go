package sse

import (
	"net/http"

	"github.com/thereisnoplanb/http/statusRecorder"
)

func New(response http.ResponseWriter) ResponseWriter {
	response.Header().Set("Content-Type", "text/event-stream")
	response.Header().Set("Cache-Control", "no-cache")
	response.Header().Set("Connection", "keep-alive")

	return &responseWriter{
		ResponseWriter: statusRecorder.New(response),
	}
}
