package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/stores/internal/domain"
)

type (
	RemoveProduct struct {
		ID string
	}

	RemoveProductHandler struct {
		stores   domain.StoreRepository
		products domain.ProductRepository
	}
)

func NewRemoveProductHandler(stores domain.StoreRepository, products domain.ProductRepository) RemoveProductHandler {
	return RemoveProductHandler{
		stores:   stores,
		products: products,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	return h.products.Delete(ctx, cmd.ID)
}
