package main

import (
	"sync"
)

var (
	registryInstance *Registry
	once             sync.Once
)

type Registry struct {
	/*
		Nested map structure
			[name]:
				[id] : SVC 1
				[id] : SVC 1
			[name]:
				[id] : SVC 2
				[id] : SVC 2
	*/

	services map[string]map[string]Service
	mu       sync.RWMutex
}

func NewRegistry() *Registry {
	once.Do(func() {
		instance := &Registry{
			services: make(map[string]map[string]Service),
			mu:       sync.RWMutex{},
		}

		registryInstance = instance
	})

	return registryInstance
}

func (r *Registry) Add(svc Service) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.services[svc.Name]; !exists {
		r.services[svc.Name] = make(map[string]Service)
	}

	r.services[svc.Name][svc.ID] = svc
}

func (r *Registry) Remove(name, id string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if nameSvc, exists := r.services[name]; exists {
		delete(nameSvc, id)

		if len(nameSvc) == 0 {
			delete(r.services, name)
		}
	}
}

func (r *Registry) Get(name string) []Service {
	r.mu.Lock()
	defer r.mu.Unlock()

	var services []Service

	if nameSvc, exists := r.services[name]; exists {
		for _, svc := range nameSvc {
			services = append(services, svc)
		}
	}

	return services
}
