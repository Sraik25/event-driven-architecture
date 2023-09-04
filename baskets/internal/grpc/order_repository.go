package grpc

import (
	"context"
	"google.golang.org/grpc"

	"github.com/Sraik25/event-driven-architecture/baskets/internal/domain"
)

type OrderRepository struct {
}

func NewOrderRepository(conn *grpc.ClientConn) *OrderRepository {
	return &OrderRepository{}
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func (o OrderRepository) Save(ctx context.Context, basket *domain.Basket) (string, error) {
	// TODO
	return "", nil
}
