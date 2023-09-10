package queries

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/stores/internal/domain"
)

type (
	GetParticipatingStores struct {
	}

	GetParticipatingStoresHandler struct {
		participatingStores domain.ParticipatingStoreRepository
	}
)

func NewGetParticipatingStoresHandler(participatingStores domain.ParticipatingStoreRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{participatingStores: participatingStores}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.Store, error) {
	return h.participatingStores.FindAll(ctx)
}
