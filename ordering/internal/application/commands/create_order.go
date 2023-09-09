package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
	"github.com/stackus/errors"

	"github.com/Sraik25/event-driven-architecture/ordering/internal/domain"
)

type CreateOrder struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []*domain.Item
}

type CreateOrderHandler struct {
	orders          domain.OrderRepository
	customers       domain.CustomerRepository
	payments        domain.PaymentRepository
	shopping        domain.ShoppingRepository
	domainPublisher ddd.EventPublisher
}

func NewCreateOrderHandler(orders domain.OrderRepository, customers domain.CustomerRepository,
	payments domain.PaymentRepository, shopping domain.ShoppingRepository, domainPublisher ddd.EventPublisher) CreateOrderHandler {
	return CreateOrderHandler{
		orders:          orders,
		customers:       customers,
		payments:        payments,
		shopping:        shopping,
		domainPublisher: domainPublisher,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := domain.CreateOrder(cmd.ID, cmd.CustomerID, cmd.PaymentID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	// authorizeCustomer
	if err = h.customers.Authorize(ctx, order.CustomerID); err != nil {
		return errors.Wrap(err, "order customers authorization")
	}

	// validatePayment
	if err = h.payments.Confirm(ctx, order.PaymentID); err != nil {
		return errors.Wrap(err, "order payment confirmation")
	}

	// scheduleShopping
	if order.ShoppingID, err = h.shopping.Create(ctx, order); err != nil {
		return errors.Wrap(err, "order shopping scheduling")
	}

	if err = h.orders.Save(ctx, order); err != nil {
		errors.Wrap(err, "Create order command")
	}

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
