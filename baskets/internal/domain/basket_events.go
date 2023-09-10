package domain

const (
	BasketStartedEvent     = "baskets.BasketStarted"
	BasketItemAddedEvent   = "baskets.BasketItemAddedEvent"
	BasketItemRemovedEvent = "baskets.BasketItemRemovedEvent"
	BasketCanceledEvent    = "baskets.BasketCanceledEvent"
	BasketCheckedOutEvent  = "baskets.BasketCheckedOutEvent"
)

type BasketStarted struct {
	CustomerID string
}

func (BasketStarted) Key() string { return BasketStartedEvent }

type BasketItemAdded struct {
	Item Item
}

func (BasketItemAdded) Key() string { return BasketItemAddedEvent }

type BasketItemRemoved struct {
	ProductID string
	Quantity  int
}

func (BasketItemRemoved) Key() string { return BasketItemRemovedEvent }

type BasketCanceled struct{}

func (BasketCanceled) Key() string { return BasketCanceledEvent }

type BasketCheckedOut struct {
	PaymentID  string
	CustomerID string
	Items      map[string]Item
}

func (BasketCheckedOut) Key() string { return BasketCheckedOutEvent }
