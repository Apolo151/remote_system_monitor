package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/Apolo151/remote_system_monitor/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MonitorServer struct {
	pb.UnimplementedMonitorServiceServer
}

func PrintMetrics(stream pb.MonitorService_StreamMetricsClient) error {
	for {
		metrics, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving metrics: %v", err)
		}
		log.Printf("CPU Usage: %.2f%%, RAM Usage: %.2f%%", metrics.CpuPercent, metrics.RamPercent)
	}
}

func main() {
	serverAddr := flag.String("server_addr", "localhost:50051", "The server address in the format of host:port")
	flag.Parse()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewMonitorServiceClient(conn)

	stream, err := client.StreamMetrics(context.Background(), &pb.MetricsRequest{})
	if err != nil {
		log.Fatalf("error calling StreamMetrics: %v", err)
	}
	PrintMetrics(stream)

	if err != nil {
		log.Fatalf("error calling StreamMetrics: %v", err)
	}
}
