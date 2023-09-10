package domain

import "github.com/stackus/errors"

var (
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customers id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type Order struct {
	ID         string
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []*Item
	Status     OrderStatus
}

func CreateOrder(ID, customerID, paymentID string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	return &Order{
		ID:         ID,
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      items,
		Status:     OrderPending,
	}, nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderPending {
		return ErrOrderCannotBeCancelled
	}

	o.Status = OrderCancelled
	return nil
}

func (o *Order) Ready() error {
	// TODO validate status
	o.Status = OrderReady
	return nil
}

func (o *Order) Complete(invoiceID string) error {
	// TODO validate invoice exists

	// TODO validate status
	o.InvoiceID = invoiceID
	o.Status = OrderCompleted
	return nil
}

func (o *Order) GetTotal() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
