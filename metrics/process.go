package metrics

import (
	"time"
)

type ProcessMetrics struct {
	When           time.Time
	CurrentState   ProcessState
	PID            int32
	UserCpu        uint64
	SystemCpu      uint64
	UserChildCpu   uint64
	SystemChildCpu uint64
}

type ProcessState uint8

const (
	Down ProcessState = iota
	Starting
	Running
	Stopping
)

type Process struct {
	Name           string
	PID            int32
	MetricsHistory []*ProcessMetrics
}

func (p *Process) CaptureMetrics() (*ProcessMetrics, error) {
	return &ProcessMetrics{time.Now(), Running, p.PID, 0, 0, 0, 0}, nil
}