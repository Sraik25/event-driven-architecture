package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
)

type AssignShoppingList struct {
	ID    string
	BotID string
}

type AssignShoppingListHandler struct {
	shoppingList domain.ShoppingListRepository
}

func NewAssignShoppingListHandler(shoppingList domain.ShoppingListRepository) AssignShoppingListHandler {
	return AssignShoppingListHandler{shoppingList: shoppingList}
}

func (h AssignShoppingListHandler) AssignShoppingList(ctx context.Context, cmd AssignShoppingList) error {
	list, err := h.shoppingList.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Assign(cmd.BotID)
	if err != nil {
		return err
	}

	return h.shoppingList.Update(ctx, list)
}
