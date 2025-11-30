package host

import (
	"channels/internal/helpers"
	"sync"
)

// Simulates a single unit of a service
type Pod interface {
	IP() string
	Health() *HealthStatus
}

type HealthStatus struct {
	Name         string
	Ok           bool
	Message      string
	Dependencies []DependencyStatus
}

type DependencyStatus struct {
	Ok bool
}

type pod struct {
	name, ip          string
	numOfDependencies int
}

func NewPod(name, ip string, numOfDependencies int) Pod {
	return &pod{
		name:              name,
		numOfDependencies: numOfDependencies,
		ip:                ip,
	}
}

func (p *pod) IP() string {
	return p.ip
}

func (p *pod) Health() *HealthStatus {
	status := helpers.MicroserviceLatency()
	statuses := make(chan DependencyStatus, p.numOfDependencies)
	var wg sync.WaitGroup

	// Using channels to retrieve dependencies status
	for range p.numOfDependencies {
		wg.Go(func() {
			statuses <- *fakeStatus()
		})
	}

	go func() {
		wg.Wait()
		close(statuses)
	}()

	dependencies := make([]DependencyStatus, p.numOfDependencies)
	ok := true
	for d := range statuses {
		dependencies = append(dependencies, d)
		ok = ok && d.Ok
	}

	if status > 4900 {
		ok = false
	}

	return &HealthStatus{
		Name:         p.name,
		Ok:           ok,
		Dependencies: dependencies,
	}
}

func fakeStatus() *DependencyStatus {
	status := helpers.MicroserviceLatency()
	hasError := status > 4900

	return &DependencyStatus{
		Ok: hasError,
	}
}
