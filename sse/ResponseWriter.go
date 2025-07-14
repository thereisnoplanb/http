package sse

import (
	"time"

	"github.com/thereisnoplanb/http/statusRecorder"
)

type ResponseWriter interface {
	statusRecorder.ResponseWriter

	// Sends a Server-Sent Event (SSE) to the client.
	//
	// # Parameters
	//
	//	id string
	//
	// A unique identifier for the event. Useful for clients to track missed events.
	//
	//	event string
	//
	// The event type (e.g., "message", "update", "error"). Helps the client distinguish between data types.
	//
	// 	data any
	//
	// Any payload associated with the event. It must be serializable to text.
	// Supported types:
	// string,
	// int, int8, int16, int32, int64,
	// uint, uint8, uint16, uint32, uint64, uintptr,
	// float32, float64,
	// complex64, complex128,
	// bool,
	// time.Time,
	// []byte.
	// Other types are sent as JSON.
	// # Returns
	//
	//	err error
	//
	// An error if something goes wrong during event streaming; nil otherwise.
	SendStreamEvent(id string, event string, data any) (err error)

	// Sends a keep-alive signal to the client over the SSE connection.
	//
	// This method is typically used to maintain the connection during periods
	// of inactivity and prevent timeouts by intermediaries.
	//
	// # Returns
	//
	// 	err error
	//
	// An error if the ping signal could not be sent; nil otherwise.
	Ping() (err error)

	// Sets the reconnection time for the client in case of connection loss.
	//
	// # Parameters
	//
	//	retryDelay time.Duration
	//
	// The duration the client should wait before attempting to reconnect (e.g., 3 * time.Second).
	//
	// # Returns
	//
	//	err error
	//
	// An error if the retry directive could not be sent; nil otherwise.
	SendRetry(retryDelay time.Duration) (err error)
}
