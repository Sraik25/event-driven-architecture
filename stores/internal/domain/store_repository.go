package domain

import "context"

type StoreRepository interface {
	Load(ctx context.Context, storeID string) (*Store, error)
	Save(ctx context.Context, store *Store) error
	/*Update(ctx context.Context, store *Store) error
	Delete(ctx context.Context, storeID string) error
	Find(ctx context.Context, storeID string) (*Store, error)
	FindAll(ctx context.Context) ([]*Store, error)*/
}
