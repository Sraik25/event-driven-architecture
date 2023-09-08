package application

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
