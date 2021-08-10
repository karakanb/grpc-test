package repository

import (
	"errors"
	"github.com/karakanb/plans-service-grpc-poc/pkg/model"
	"time"
)

type PlansRepository struct{}

func NewPlansRepository() *PlansRepository {
	return &PlansRepository{}
}

func (p PlansRepository) GetPlan(ID string) (*model.Plan, error) {
	if ID == "error" {
		return nil, errors.New("something bad happened while getting the plan")
	}

	return &model.Plan{
		ID:           ID,
		CustomerUUID: "some customer ID",
		Sku:          "some sku",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
