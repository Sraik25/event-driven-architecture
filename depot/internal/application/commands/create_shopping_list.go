package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
	"github.com/stackus/errors"

	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
)

type CreateShoppingList struct {
	ID      string
	OrderID string
	Items   []OrderItem
}

type CreateShoppingListHandler struct {
	shoppingList    domain.ShoppingListRepository
	stores          domain.StoreRepository
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher
}

func NewCreateShoppingListHandler(shoppingList domain.ShoppingListRepository, stores domain.StoreRepository,
	products domain.ProductRepository, domainPublisher ddd.EventPublisher) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingList:    shoppingList,
		stores:          stores,
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h CreateShoppingListHandler) CreateShoppingList(ctx context.Context, cmd CreateShoppingList) error {
	list := domain.CreateShopping(cmd.ID, cmd.OrderID)

	for _, item := range cmd.Items {
		store, err := h.stores.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}

		product, err := h.products.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}

		err = list.AddItem(store, product, item.Quantity)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
	}

	if err := h.shoppingList.Save(ctx, list); err != nil {
		return errors.Wrap(err, "scheduling shopping")
	}

	// publish domain events
	if err := h.domainPublisher.Publish(ctx, list.GetEvents()...); err != nil {
		return err
	}

	return nil
}
