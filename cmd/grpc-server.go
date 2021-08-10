package main

import (
	"fmt"
	"github.com/karakanb/plans-service-grpc-poc/pkg/handler"
	"github.com/karakanb/plans-service-grpc-poc/pkg/repository"
	plansservice "github.com/karakanb/plans-service-grpc-poc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	GRPC_PORT = 8000
)

func main() {
	grpcListen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	plansRepo := repository.NewPlansRepository()
	plansservice.RegisterPlansServiceServer(grpcServer, handler.NewGrpcServer(handler.NewHandler(plansRepo)))

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown", s)
		grpcServer.GracefulStop()
		wg.Done()
	}()

	log.Printf("running the server on port %d\n", GRPC_PORT)
	err = grpcServer.Serve(grpcListen)
	if err != nil {
		log.Fatalf("failed to serve the grpc server: %v", err)
	}

	wg.Wait()
	log.Println("clean shutdown")
}
