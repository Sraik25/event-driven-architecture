package handlers

import (
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/application"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
