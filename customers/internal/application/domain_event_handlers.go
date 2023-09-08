package application

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

type DomainEventHandlers interface {
	OnCustomerRegistered(ctx context.Context, event ddd.Event) error
	OnCustomerAuthorized(ctx context.Context, event ddd.Event) error
	OnCustomerEnabled(ctx context.Context, event ddd.Event) error
	OnCustomerDisabled(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (i ignoreUnimplementedDomainEvents) OnCustomerRegistered(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnCustomerAuthorized(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnCustomerEnabled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnCustomerDisabled(ctx context.Context, event ddd.Event) error {
	return nil
}
