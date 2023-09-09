package application

import (
	"context"

	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

type DomainEventHandlers interface {
	OnOrderCreated(ctx context.Context, event ddd.Event) error
	OnOrderReadied(ctx context.Context, event ddd.Event) error
	OnOrderCanceled(ctx context.Context, event ddd.Event) error
	OnOrderCompleted(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (i ignoreUnimplementedDomainEvents) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnOrderReadied(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnOrderCanceled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnOrderCompleted(ctx context.Context, event ddd.Event) error {
	return nil
}
