package sse

import "github.com/thereisnoplanb/http/statusRecorder"

type ResponseWriter interface {
	statusRecorder.ResponseWriter
	SendStreamEvent(id string, event string, data []byte) (err error)
}
