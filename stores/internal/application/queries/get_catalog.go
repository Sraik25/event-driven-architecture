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
		catalog domain.CatalogRepository
	}
)

func NewGetCatalogHandler(catalog domain.CatalogRepository) GetCatalogHandler {
	return GetCatalogHandler{
		catalog: catalog,
	}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*domain.CatalogProduct, error) {
	return h.catalog.GetCatalog(ctx, query.StoreID)
}
