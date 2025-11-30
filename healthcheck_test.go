package main

import (
	"channels/internal/host"
	"sync"
	"testing"
	"time"
)

var serviceNames []string = []string{
	"Authentication",
	"Authorization",
	"Account",
	"Gateway",
	"Store",
	"Leaderboard",
	"Challenges",
	"Matchmaking",
	"Friends",
	"Population",
	"Avatars",
	"Profiles",
	"Communication",
}

var cluster host.Cluster
var services []host.Microservice

func init() {
	cluster = host.NewCluster(serviceNames)
	services = cluster.GetServices()
}

func TestHealthCheckSync(t *testing.T) {
	start := time.Now()
	for _, s := range services {
		s.GetHealth()
	}

	t.Logf("(%f s) Test completer after", time.Since(start).Seconds())
}

func TestHealthCheckAsync(t *testing.T) {
	healthChecks := make(chan *host.HealthStatus, len(services))
	durations := make(chan time.Duration, len(services))
	start := time.Now()
	var wg sync.WaitGroup
	for _, s := range services {
		wg.Go(func() {
			health, duration := s.GetHealth()
			healthChecks <- health
			durations <- duration
		})
	}

	go func() {
		wg.Wait()
		close(healthChecks)
		close(durations)
	}()

	var receivedChecks int
	minDuration := time.Hour
	maxDuration := time.Nanosecond
	for range healthChecks {
		receivedChecks++
	}

	for d := range durations {
		if d > maxDuration {
			maxDuration = d
		}
		if d < minDuration {
			minDuration = d
		}
	}

	t.Logf("(%f s) %d healthchecks received of %d - min %f s and max %f s \n", time.Since(start).Seconds(), receivedChecks, len(services), minDuration.Seconds(), maxDuration.Seconds())
}

func TestHealthCheckAsyncWithTimeout(t *testing.T) {
	timeout := 10 * time.Second
	expectedCount := len(services)
	healthChecks := make(chan *host.HealthStatus, expectedCount)
	durations := make(chan time.Duration, expectedCount)
	start := time.Now()
	var wg sync.WaitGroup

	for _, s := range services {
		wg.Go(func() {
			workerHealth := make(chan *host.HealthStatus)
			workerDuration := make(chan time.Duration)

			go func() {
				health, duration := s.GetHealth()
				workerHealth <- health
				workerDuration <- duration
			}()
			select {
			case <-time.After(timeout):
				t.Logf("Warning: Health check for service %s took longer than %s and was cancelled.", s.GetName(), timeout)
			case health := <-workerHealth:
				duration := <-workerDuration
				healthChecks <- health
				durations <- duration
			}
		})
	}

	go func() {
		wg.Wait()
		close(healthChecks)
		close(durations)
	}()

	var receivedChecks int
	minDuration := time.Hour
	maxDuration := time.Nanosecond
	for range healthChecks {
		receivedChecks++
	}

	for d := range durations {
		if d > maxDuration {
			maxDuration = d
		}
		if d < minDuration {
			minDuration = d
		}
	}

	t.Logf("(%f s): %d healthchecks received of %d - min %f s and max %f s",
		time.Since(start).Seconds(), receivedChecks, expectedCount, minDuration.Seconds(), maxDuration.Seconds())
}
