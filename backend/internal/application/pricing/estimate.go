package pricingapp

import (
	"context"
	"errors"
	"math"

	domainpricing "github.com/armelo10/vtc_go/backend/internal/domain/pricing"
)

type RoutingProvider interface {
	Estimate(context.Context, float64, float64, float64, float64) (float64, float64, error)
}

type Service struct {
	routing RoutingProvider
	engine  domainpricing.Engine
}

func NewService(routing RoutingProvider, engine domainpricing.Engine) *Service {
	return &Service{routing: routing, engine: engine}
}

func (s *Service) Estimate(ctx context.Context, a, b, c, d float64) (domainpricing.Quote, float64, float64, error) {
	if s.routing == nil {
		return domainpricing.Quote{}, 0, 0, errors.New("routing provider unavailable")
	}
	km, min, err := s.routing.Estimate(ctx, a, b, c, d)
	if err != nil {
		return domainpricing.Quote{}, 0, 0, err
	}
	if km <= 0 || min <= 0 || math.IsNaN(km) || math.IsNaN(min) {
		return domainpricing.Quote{}, 0, 0, errors.New("invalid routing estimate")
	}
	return s.engine.Estimate(km, min), km, min, nil
}
