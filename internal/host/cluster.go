package host

import (
	"channels/internal/helpers"
	"fmt"
)

// Simulate a cluster where many microservices are alive
type Cluster interface {
	GetServices() []Microservice
}

type cluster struct {
	services []Microservice
}

func NewCluster(serviceNames []string) Cluster {
	services := make([]Microservice, len(serviceNames))

	for i, s := range serviceNames {
		deps := helpers.Random(1, 10)
		numOfPods := helpers.Random(1, 100)
		ips := generateIps(i, numOfPods)
		services[i] = NewMicroService(s, ips, deps)
	}

	return &cluster{
		services: services,
	}
}

func (c *cluster) GetServices() []Microservice {
	helpers.NetworkLatency()
	return c.services
}

func generateIps(index, numOfIps int) []string {
	base := "192.168.%d.%d"
	ips := make([]string, numOfIps)

	for i := range numOfIps {
		ips[i] = fmt.Sprintf(base, index, i)
	}

	return ips
}
