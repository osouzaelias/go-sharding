package api

import (
	"context"
	"fmt"
	"github.com/cespare/xxhash/v2"
	"go-sharding/internal/application/core/domain"
	"go-sharding/internal/ports"
)

type Application struct {
	db         ports.DBPort
	nodes      []string
	rendezvous *domain.Rendezvous
}

func NewApplication(db ports.DBPort, nodes []string) *Application {
	return &Application{
		db:         db,
		nodes:      nodes,
		rendezvous: domain.NewRendezvous(nodes, xxhash.Sum64String),
	}
}

func (a Application) Add(ctx context.Context, item domain.Customer) error {
	node := a.rendezvous.Lookup(item.ID)

	if err := a.db.Create(ctx, item, node); err != nil {
		return err
	}

	fmt.Println("record stored on", node)

	return nil
}

func (a Application) Get(ctx context.Context, customerID string) (domain.Customer, error) {
	node := a.rendezvous.Lookup(customerID)

	customer, err := a.db.Read(ctx, customerID, node)
	if err != nil {
		return domain.Customer{}, nil
	}

	fmt.Println("record retrieved from", node)

	return customer, nil
}
