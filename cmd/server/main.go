package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Apolo151/remote_system_monitor/internals/config"
	"github.com/Apolo151/remote_system_monitor/internals/server"
	pb "github.com/Apolo151/remote_system_monitor/pkg/monitorpb"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadServerConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMonitorServiceServer(grpcServer, server.NewMonitorServer(cfg.Interval))

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down server...")
		grpcServer.GracefulStop()
	}()

	log.Printf("Server listening on :%d (interval: %v)", cfg.Port, cfg.Interval)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
