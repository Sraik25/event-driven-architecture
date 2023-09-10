package queries

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
)

type GetShoppingList struct {
	ID string
}

type GetShoppingListHandler struct {
	shoppingList domain.ShoppingListRepository
}

func NewGetShoppingListHandler(shoppingList domain.ShoppingListRepository) GetShoppingListHandler {
	return GetShoppingListHandler{shoppingList: shoppingList}
}

func (h GetShoppingListHandler) GetShoppingList(ctx context.Context, query GetShoppingList) (*domain.ShoppingList, error) {
	return h.shoppingList.Find(ctx, query.ID)
}
