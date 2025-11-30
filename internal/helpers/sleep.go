package helpers

import (
	"time"
)

// Calculate a time to sleep between 0 and duration to sleep and return this random number
// Function to simulate a request call and return on a response
func NetworkLatency() int {
	miliseconds := Random(100, 500) // variation between 100ms and 500ms
	time.Sleep(time.Duration(miliseconds) * time.Millisecond)
	return miliseconds
}

func MicroserviceLatency() int {
	miliseconds := Random(500, 10*1000) // variation between 500ms and 10s
	time.Sleep(time.Duration(miliseconds) * time.Millisecond)
	return int(miliseconds)
}
