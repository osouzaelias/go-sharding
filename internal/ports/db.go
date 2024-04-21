package ports

import (
	"context"
	"go-sharding/internal/application/core/domain"
)

type DBPort interface {
	Create(ctx context.Context, customer domain.Customer, node string) error
	Read(ctx context.Context, customerID, node string) (domain.Customer, error)
}
