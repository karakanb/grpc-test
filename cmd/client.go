package main

import (
	"context"
	"flag"
	"fmt"
	plansservice "github.com/karakanb/plans-service-grpc-poc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

func main() {
	flag.Parse()

	host := fmt.Sprintf("localhost:%d", *port)
	fmt.Println("Running")

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := plansservice.NewPlansServiceClient(conn)

	for true {
		plan, err := client.GetPlan(context.Background(), &plansservice.GetPlanRequest{ID: "123123"})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Fetched a plan: %v\n", plan)
		time.Sleep(time.Second * 3)
	}
}
