package monitor

import "context"

type Monitor struct {
}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) Inspect(_ context.Context) error {
	return nil
}

func (m *Monitor) Shutdown() error {
	return nil
}
