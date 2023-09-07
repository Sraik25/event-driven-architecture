package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
)

type CancelShoppingList struct {
	ID string
}

type CancelShoppingListHandler struct {
	shoppingList domain.ShoppingListRepository
}

func NewCancelShoppingListHandler(shoppingList domain.ShoppingListRepository) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingList: shoppingList,
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

	return h.shoppingList.Update(ctx, list)
}
