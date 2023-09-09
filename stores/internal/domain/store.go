package domain

import (
	"github.com/stackus/errors"

	"github.com/Sraik25/event-driven-architecture/internal/es"
)

const StoreAggregate = "stores.Store"

var (
	ErrStoreNameIsBlank               = errors.Wrap(errors.ErrBadRequest, "the store name cannot be blank")
	ErrStoreLocationIsBlank           = errors.Wrap(errors.ErrBadRequest, "the store location cannot be blank")
	ErrStoreIsAlreadyParticipating    = errors.Wrap(errors.ErrBadRequest, "the store is already participating")
	ErrStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "the store is already not participating")
)

type Store struct {
	es.Aggregate
	Name          string
	Location      string
	Participating bool
}

func NewStore(id string) *Store {
	return &Store{
		Aggregate: es.NewAggregate(id, StoreAggregate),
	}
}

func CreateStore(id, name, location string) (store *Store, err error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

	store = NewStore(id)

	store.AddEvent(StoreCreatedEvent, &StoreCreated{
		Store: store,
	})

	return
}

func (s *Store) EnableParticipation() (err error) {
	if s.Participating {
		return ErrStoreIsAlreadyParticipating
	}

	s.Participating = true

	s.AddEvent(StoreParticipationEnabledEvent, &StoreParticipationEnabled{
		Store: s,
	})

	return
}

func (s *Store) DisableParticipation() (err error) {
	if !s.Participating {
		return ErrStoreIsAlreadyNotParticipating
	}

	s.Participating = false

	s.AddEvent(StoreParticipationDisabledEvent, &StoreParticipationDisabled{
		Store: s,
	})

	return
}
