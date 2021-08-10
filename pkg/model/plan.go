package model

import (
	plansservice "github.com/karakanb/plans-service-grpc-poc/proto"
	"time"
)

type Plan struct {
	ID           string
	CustomerUUID string
	Sku          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (p Plan) ToProtobuf() *plansservice.Plan {
	return &plansservice.Plan{
		ID:           p.ID,
		CustomerUUID: p.CustomerUUID,
		Sku:          p.Sku,
		CreatedAt:    p.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:    p.UpdatedAt.Format(time.RFC3339Nano),
	}
}
