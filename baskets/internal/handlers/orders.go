package handlers

import (
	"github.com/Sraik25/event-driven-architecture/baskets/internal/domain"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.BasketCheckedOutEvent, orderHandlers)
}
