package commands

import (
	"context"
	"github.com/stackus/errors"

	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
)

type CreateShoppingList struct {
	ID      string
	OrderID string
	Items   []OrderItem
}

type CreateShoppingListHandler struct {
	shoppingList domain.ShoppingListRepository
	stores       domain.StoreRepository
	products     domain.ProductRepository
}

func NewCreateShoppingListHandler(shoppingList domain.ShoppingListRepository, stores domain.StoreRepository, products domain.ProductRepository) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingList: shoppingList,
		stores:       stores,
		products:     products,
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

	return errors.Wrap(h.shoppingList.Save(ctx, list), "scheduling shopping")
}
