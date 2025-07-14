package sse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/thereisnoplanb/http/statusRecorder"
)

type responseWriter struct {
	statusRecorder.ResponseWriter
}

func (response *responseWriter) SendStreamEvent(id string, event string, data any) (err error) {
	switch value := data.(type) {
	case string:
		return response.sendString(id, event, value)
	case int:
		return response.sendInt(id, event, value)
	case int8:
		return response.sendInt8(id, event, value)
	case int16:
		return response.sendInt16(id, event, value)
	case int32:
		return response.sendInt32(id, event, value)
	case int64:
		return response.sendInt64(id, event, value)
	case uint:
		return response.sendUInt(id, event, value)
	case uint8:
		return response.sendUInt8(id, event, value)
	case uint16:
		return response.sendUInt16(id, event, value)
	case uint32:
		return response.sendUInt32(id, event, value)
	case uint64:
		return response.sendUInt64(id, event, value)
	case float32:
		return response.sendFloat32(id, event, value)
	case float64:
		return response.sendFloat64(id, event, value)
	case complex64:
		return response.sendComplex64(id, event, value)
	case complex128:
		return response.sendComplex128(id, event, value)
	case uintptr:
		return response.sendUIntPtr(id, event, value)
	case bool:
		return response.sendBool(id, event, value)
	case time.Time:
		return response.sendTime(id, event, value)
	case []byte:
		return response.sendBytes(id, event, value)
	default:
		return response.sendJSON(id, event, value)
	}
}

func (response *responseWriter) sendString(id string, event string, value string) (err error) {
	return response.send(id, event, value)
}

func (response *responseWriter) sendInt(id string, event string, value int) (err error) {
	return response.send(id, event, strconv.Itoa(value))
}

func (response *responseWriter) sendInt8(id string, event string, value int8) (err error) {
	return response.send(id, event, strconv.FormatInt(int64(value), 10))
}

func (response *responseWriter) sendInt16(id string, event string, value int16) (err error) {
	return response.send(id, event, strconv.FormatInt(int64(value), 10))
}

func (response *responseWriter) sendInt32(id string, event string, value int32) (err error) {
	return response.send(id, event, strconv.FormatInt(int64(value), 10))
}

func (response *responseWriter) sendInt64(id string, event string, value int64) (err error) {
	return response.send(id, event, strconv.FormatInt(value, 10))
}

func (response *responseWriter) sendUInt(id string, event string, value uint) (err error) {
	return response.send(id, event, strconv.FormatUint(uint64(value), 10))
}

func (response *responseWriter) sendUInt8(id string, event string, value uint8) (err error) {
	return response.send(id, event, strconv.FormatUint(uint64(value), 10))
}

func (response *responseWriter) sendUInt16(id string, event string, value uint16) (err error) {
	return response.send(id, event, strconv.FormatUint(uint64(value), 10))
}

func (response *responseWriter) sendUInt32(id string, event string, value uint32) (err error) {
	return response.send(id, event, strconv.FormatUint(uint64(value), 10))
}

func (response *responseWriter) sendUInt64(id string, event string, value uint64) (err error) {
	return response.send(id, event, strconv.FormatUint(value, 10))
}

func (response *responseWriter) sendFloat32(id string, event string, value float32) (err error) {
	return response.send(id, event, strconv.FormatFloat(float64(value), 'f', -1, 32))
}

func (response *responseWriter) sendFloat64(id string, event string, value float64) (err error) {
	return response.send(id, event, strconv.FormatFloat(value, 'f', -1, 64))
}

func (response *responseWriter) sendComplex64(id string, event string, value complex64) (err error) {
	return response.send(id, event, strconv.FormatComplex(complex128(value), 'f', -1, 64))
}

func (response *responseWriter) sendComplex128(id string, event string, value complex128) (err error) {
	return response.send(id, event, strconv.FormatComplex(value, 'f', -1, 128))
}

func (response *responseWriter) sendUIntPtr(id string, event string, value uintptr) (err error) {
	return response.send(id, event, strconv.FormatUint(uint64(value), 10))
}

func (response *responseWriter) sendBool(id string, event string, value bool) (err error) {
	return response.send(id, event, strconv.FormatBool(value))
}

func (response *responseWriter) sendTime(id string, event string, value time.Time) (err error) {
	return response.send(id, event, value.String())
}

func (response *responseWriter) sendBytes(id string, event string, value []byte) (err error) {
	return response.send(id, event, string(value))
}

func (response *responseWriter) sendJSON(id string, event string, object any) (err error) {
	data, err := json.Marshal(object)
	if err != nil {
		return err
	}
	return response.send(id, event, string(data))
}

func (response *responseWriter) send(id string, event string, value string) (err error) {
	buffer := &bytes.Buffer{}
	_, _ = fmt.Fprintf(buffer, "id: %s\n", id)
	_, _ = fmt.Fprintf(buffer, "event: %s\n", event)
	for value = range strings.SplitSeq(value, "\n") {
		_, _ = fmt.Fprintf(buffer, "data: %s\n", value)
	}
	_, _ = buffer.WriteRune('\n')
	return response.sendBuffer(buffer)
}

func (response *responseWriter) Ping() (err error) {
	return response.ping()
}

func (response *responseWriter) ping() (err error) {
	buffer := &bytes.Buffer{}
	_, _ = fmt.Fprint(buffer, ": ping\n\n")
	return response.sendBuffer(buffer)
}

func (response *responseWriter) SendRetry(retryDelay time.Duration) (err error) {
	return response.retry(retryDelay)
}

func (response *responseWriter) retry(retryDelay time.Duration) (err error) {
	buffer := &bytes.Buffer{}
	_, _ = fmt.Fprintf(buffer, "retry: %d\n\n", int(retryDelay/time.Millisecond))
	return response.sendBuffer(buffer)
}

func (response *responseWriter) sendBuffer(buffer *bytes.Buffer) (err error) {
	_, err = buffer.WriteTo(response)
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
