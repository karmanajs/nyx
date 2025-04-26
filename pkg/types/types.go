package types

import "time"

// Config holds application configuration
type Config struct {
	Host       string
	Ports      string
	Protocol   string
	OutputJSON string
	Timeout    time.Duration
}

// ScanResult represents the result of a port scan
type ScanResult struct {
	Port    int    `json:"port"`
	Status  string `json:"status"`
	Service string `json:"service,omitempty"`
}
