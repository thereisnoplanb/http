package statusRecorder

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (response *responseWriter) WriteHeader(statusCode int) {
	response.statusCode = statusCode
	response.ResponseWriter.WriteHeader(statusCode)
}

func (response *responseWriter) StatusCode() int {
	return response.statusCode
}

func (response *responseWriter) Flush() {
	if flusher, ok := response.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}
