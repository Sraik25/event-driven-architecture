package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

type CancelShoppingList struct {
	ID string
}

type CancelShoppingListHandler struct {
	shoppingList    domain.ShoppingListRepository
	domainPublisher ddd.EventPublisher
}

func NewCancelShoppingListHandler(shoppingList domain.ShoppingListRepository, domainPublisher ddd.EventPublisher) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingList:    shoppingList,
		domainPublisher: domainPublisher,
	}
}

func (h CancelShoppingListHandler) CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error {
	list, err := h.shoppingList.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Cancel()
	if err != nil {
		return err
	}

	if err = h.shoppingList.Update(ctx, list); err != nil {
		return err
	}

	// publish domain events
	if err := h.domainPublisher.Publish(ctx, list.GetEvents()...); err != nil {
		return err
	}

	return nil
}
