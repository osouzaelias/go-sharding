package ports

import (
	"context"
	"go-sharding/internal/application/core/domain"
)

type APIPort interface {
	Add(ctx context.Context, customer domain.Customer) error
	Get(ctx context.Context, customerID string) (domain.Customer, error)
}
