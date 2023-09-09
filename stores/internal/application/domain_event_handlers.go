package application

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

type DomainEventHandlers interface {
	OnStoreCreated(ctx context.Context, event ddd.Event) error
	OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error
	OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error
	OnProductAdded(ctx context.Context, event ddd.Event) error
	OnProductRemoved(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (i ignoreUnimplementedDomainEvents) OnStoreCreated(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnProductAdded(ctx context.Context, event ddd.Event) error {
	return nil
}

func (i ignoreUnimplementedDomainEvents) OnProductRemoved(ctx context.Context, event ddd.Event) error {
	return nil
}
