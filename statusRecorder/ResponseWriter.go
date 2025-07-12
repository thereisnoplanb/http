package statusRecorder

import "net/http"

type ResponseWriter interface {
	http.ResponseWriter
	StatusCode() int
}
