package domain

const (
	StoreCreatedEvent               = "stores.StoreCreated"
	StoreParticipationEnabledEvent  = "stores.StoreParticipationEnabled"
	StoreParticipationDisabledEvent = "stores.StoreParticipationDisabled"
)

type StoreCreated struct {
	Name     string
	Location string
}

type StoreParticipationEnabled struct {
	Store *Store
}

type StoreParticipationDisabled struct {
	Store *Store
}
