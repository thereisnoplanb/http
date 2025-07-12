package statusRecorder

import "net/http"

func New(response http.ResponseWriter) ResponseWriter {
	return &responseWriter{
		ResponseWriter: response,
		statusCode:     http.StatusOK,
	}
}
