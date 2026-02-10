package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Apolo151/remote_system_monitor/internals/config"
	pb "github.com/Apolo151/remote_system_monitor/pkg/monitorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.LoadClientConfig()

	conn, err := grpc.Dial(cfg.ServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMonitorServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Shutting down client...")
		cancel()
	}()

	stream, err := client.StreamMetrics(ctx, &pb.MetricsRequest{})
	if err != nil {
		log.Fatalf("Failed to start stream: %v", err)
	}

	log.Printf("Connected to server at %s", cfg.ServerAddr)
	for {
		metrics, err := stream.Recv()
		if err != nil {
			log.Printf("Stream closed: %v", err)
			return
		}

		fmt.Printf("[%s] CPU: %.2f%% | RAM: %.2f%% (%.2f GB / %.2f GB)\n",
			time.Now().Format("15:04:05"),
			metrics.CpuPercent,
			metrics.RamPercent,
			float64(metrics.RamUsed)/1024/1024/1024,
			float64(metrics.RamTotal)/1024/1024/1024,
		)
	}
}
