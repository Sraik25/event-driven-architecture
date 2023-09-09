package handlers

import (
	"github.com/Sraik25/event-driven-architecture/depot/internal/application"
	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompleted{}, orderHandlers.OnShoppingListCompleted)
}
