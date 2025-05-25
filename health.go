package main

import "sync"

type HealthCheckManager struct {
	registry *Registry
	mu       sync.RWMutex
}

func NewHealthCheckManager(registry *Registry) *HealthCheckManager {
	return &HealthCheckManager{
		registry: registry,
		mu:       sync.RWMutex{},
	}
}

func (h *HealthCheckManager) CheckHealth(service Service) {

}
