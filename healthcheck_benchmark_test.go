package main

import (
	"channels/internal/host"
	"testing"
)

func init() {
	cluster = host.NewCluster(serviceNames)
	services = cluster.GetServices()
}

func BenchmarkHealthCheckSync(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		for _, s := range services {
			s.GetHealth()
		}
	}
}
