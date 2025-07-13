package sse

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/thereisnoplanb/http/statusRecorder"
)

type responseWriter struct {
	statusRecorder.ResponseWriter
}

func (response *responseWriter) SendStreamEvent(id string, event string, data []byte) (err error) {
	bufer := &bytes.Buffer{}
	_, _ = bufer.WriteString(fmt.Sprintf("id: %s\n", id))
	_, _ = bufer.WriteString(fmt.Sprintf("event: %s\n", event))
	_, _ = bufer.WriteString(fmt.Sprintf("data: %s\n", string(data)))
	_, _ = bufer.WriteRune('\n')

	_, err = bufer.WriteTo(response)
	if err != nil {
		return err
	}
	if flusher, ok := response.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	} else {
		return errors.New("invalid flusher")
	}
	return nil
}
