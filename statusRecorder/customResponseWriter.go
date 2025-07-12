package statusRecorder

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (responseWriter *responseWriter) WriteHeader(statusCode int) {
	responseWriter.statusCode = statusCode
	responseWriter.ResponseWriter.WriteHeader(statusCode)
}

func (responseWriter *responseWriter) StatusCode() int {
	return responseWriter.statusCode
}
