package handlers

import (
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, invoiceHandlers)
}
