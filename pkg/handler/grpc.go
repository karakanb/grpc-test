package handler

import (
	"context"
	plansservice "github.com/karakanb/plans-service-grpc-poc/proto"
)

type grpcServer struct {
	plansservice.UnimplementedPlansServiceServer

	handler *handler
}

func NewGrpcServer(handler *handler) *grpcServer {
	return &grpcServer{handler: handler}
}

func (s *grpcServer) CreatePlan(ctx context.Context, req *plansservice.PlanRequest) (*plansservice.PlanCreated, error) {
	return s.handler.CreatePlan(ctx, req)
}

func (s *grpcServer) GetPlan(ctx context.Context, req *plansservice.GetPlanRequest) (*plansservice.Plan, error) {
	return s.handler.GetPlan(ctx, req)
}
