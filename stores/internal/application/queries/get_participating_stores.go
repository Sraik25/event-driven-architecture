package queries

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/stores/internal/domain"
)

type (
	GetParticipatingStores struct {
	}

	GetParticipatingStoresHandler struct {
		mall domain.MallRepository
	}
)

func NewGetParticipatingStoresHandler(mall domain.MallRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{
		mall: mall,
	}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.MallStore, error) {
	return h.mall.AllParticipating(ctx)
}
