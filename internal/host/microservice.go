package host

import (
	"channels/internal/helpers"
	"time"
)

// Simulates a microservice with many pods and a health check endpoint with many dependencies
type Microservice interface {
	GetName() string
	GetPods() []Pod
	GetHealth() (*HealthStatus, time.Duration)
}

type microservice struct {
	name string
	pods []Pod
}

func NewMicroService(name string, ips []string, numOfDeps int) Microservice {
	pods := make([]Pod, 0)
	for _, ip := range ips {
		pods = append(pods, NewPod(name, ip, numOfDeps))
	}

	return &microservice{
		name: name,
		pods: pods,
	}
}

func (m *microservice) GetName() string {
	helpers.NetworkLatency()
	return m.name
}

func (m *microservice) GetPods() []Pod {
	helpers.NetworkLatency()
	return m.pods
}

func (m *microservice) GetHealth() (*HealthStatus, time.Duration) {
	var r int
	if len(m.pods) == 1 {
		r = 0
	} else {
		r = helpers.Random(1, len(m.pods)-1)
	}

	start := time.Now()

	status := m.pods[r].Health()

	return status, time.Since(start)
}
