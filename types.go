package main

import "time"

/*
Necessary data for registring service:
1. ID - Unique Identifier
2. Name
3. Tags []string
4. Address
5. Port
6. Metadata map[string]string
7. Health Check Config
*/
type Service struct {
	ID           string            `json:","`
	Name         string            `json:","`
	Tags         []string          `json:","`
	Port         int               `json:","`
	Address      string            `json:","`
	HealthConfig HealthCheckConfig `json:","`
	LastChecked  time.Time         `json:","`
	IsHealthy    bool              `json:","`
}

type HealthCheckConfig struct {
	URL             string `json:","`
	IntervalSeconds int    `json:","` // Seconds
	Timeout         int    `json:","` // Seconds
}
