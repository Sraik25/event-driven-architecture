package handlers

import (
	"github.com/Sraik25/event-driven-architecture/baskets/internal/application"
	"github.com/Sraik25/event-driven-architecture/baskets/internal/domain"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
