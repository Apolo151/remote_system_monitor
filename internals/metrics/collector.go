package metrics

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Collector struct {
	interval time.Duration
}

func NewCollector(interval time.Duration) *Collector {
	return &Collector{interval: interval}
}

type SystemMetrics struct {
	CPUPercent float64
	RAMPercent float64
	RAMTotal   uint64
	RAMUsed    uint64
	Timestamp  time.Time
}

func (c *Collector) Collect(ctx context.Context) (*SystemMetrics, error) {
	// CPU usage with timeout
	percentages, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil {
		return nil, err
	}

	cpuPercent := 0.0
	if len(percentages) > 0 {
		cpuPercent = percentages[0]
	}

	// RAM usage
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return &SystemMetrics{
		CPUPercent: cpuPercent,
		RAMPercent: vmStat.UsedPercent,
		RAMTotal:   vmStat.Total,
		RAMUsed:    vmStat.Used,
		Timestamp:  time.Now(),
	}, nil
}
