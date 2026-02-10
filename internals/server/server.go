package server

import (
	"log"
	"time"

	"github.com/Apolo151/remote_system_monitor/internals/metrics"
	pb "github.com/Apolo151/remote_system_monitor/pkg/monitorpb"
)

type MonitorServer struct {
	pb.UnimplementedMonitorServiceServer
	collector *metrics.Collector
	interval  time.Duration
}

func NewMonitorServer(interval time.Duration) *MonitorServer {
	return &MonitorServer{
		collector: metrics.NewCollector(interval),
		interval:  interval,
	}
}

func (s *MonitorServer) StreamMetrics(req *pb.MetricsRequest, stream pb.MonitorService_StreamMetricsServer) error {
	ctx := stream.Context()
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		case <-ticker.C:
			metrics, err := s.collector.Collect(ctx)
			if err != nil {
				log.Printf("Error collecting metrics: %v", err)
				continue
			}

			pbMetrics := &pb.SystemMetrics{
				CpuPercent: metrics.CPUPercent,
				RamPercent: metrics.RAMPercent,
				RamTotal:   metrics.RAMTotal,
				RamUsed:    metrics.RAMUsed,
			}

			if err := stream.Send(pbMetrics); err != nil {
				log.Printf("Error sending metrics: %v", err)
				return err
			}
		}
	}
}
