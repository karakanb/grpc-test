package handler

import (
	"context"
	"fmt"
	"github.com/karakanb/plans-service-grpc-poc/pkg/model"
	plansservice "github.com/karakanb/plans-service-grpc-poc/proto"
)

type PlansRepo interface {
	GetPlan(ID string) (*model.Plan, error)
}

type handler struct {
	plansRepo PlansRepo
}

func (s *handler) CreatePlan(context.Context, *plansservice.PlanRequest) (*plansservice.PlanCreated, error) {
	return &plansservice.PlanCreated{Success: true}, nil
}

func (s *handler) GetPlan(ctx context.Context, request *plansservice.GetPlanRequest) (*plansservice.Plan, error) {
	plan, err := s.plansRepo.GetPlan(request.GetID())
	if err != nil {
		return nil, fmt.Errorf("cannot get plan %s: %w", request.GetID(), err)
	}

	return plan.ToProtobuf(), nil
}

func NewHandler(repo PlansRepo) *handler {
	return &handler{
		plansRepo: repo,
	}
}
