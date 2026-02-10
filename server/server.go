package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Apolo151/remote_system_monitor/monitor"
	"google.golang.org/grpc"
)

const port = 50051

type MonitorServer struct {
	pb.UnimplementedMonitorServiceServer
}

func (s *MonitorServer) StreamMetrics(req *pb.MetricsRequest, stream pb.MonitorService_StreamMetricsServer) error {
	// This is where you would implement the logic to gather system metrics and send them to the client. For demonstration purposes, we'll just send dummy data.
	for {
		metrics := &pb.SystemMetrics{
			CpuPercent: 50.0,                    // Dummy CPU usage
			RamPercent: 70.0,                    // Dummy RAM usage
			RamTotal:   16 * 1024 * 1024 * 1024, // Dummy total RAM (16 GB)
			RamUsed:    11 * 1024 * 1024 * 1024, // Dummy used RAM (11 GB)
		}
		if err := stream.Send(metrics); err != nil {
			return err
		}
		// Sleep for a while before sending the next update
		time.Sleep(5 * time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMonitorServiceServer(grpcServer, &MonitorServer{})
	log.Printf("Server is listening on localhost:%d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
