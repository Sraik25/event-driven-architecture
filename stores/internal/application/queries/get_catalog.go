package queries

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/stores/internal/domain"
)

type (
	GetCatalog struct {
		StoreID string
	}

	GetCatalogHandler struct {
		products domain.ProductRepository
	}
)

func NewGetCatalogHandler(products domain.ProductRepository) GetCatalogHandler {
	return GetCatalogHandler{products: products}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*domain.Product, error) {
	return h.products.GetCatalog(ctx, query.StoreID)
}
