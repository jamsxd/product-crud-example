package domain

import "context"

type ProductRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindBySku(ctx context.Context, sku string) (*Product, error)
	Upsert(ctx context.Context, product *Product) error
	Delete(ctx context.Context, sku string) error
}
