package domain

import (
	"context"

	"github.com/go-kit/kit/log"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]*Product, error)
	GetProduct(ctx context.Context, sku string) (*Product, error)
	UpsertProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, sku string) error
}

type basicProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository, logger log.Logger) ProductService {
	return newLoggingMiddleware(logger)(&basicProductService{repo})
}

func (s *basicProductService) GetAllProducts(ctx context.Context) ([]*Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *basicProductService) GetProduct(ctx context.Context, sku string) (*Product, error) {
	return s.repo.FindBySku(ctx, sku)
}

func (s *basicProductService) UpsertProduct(ctx context.Context, product *Product) error {
	return s.repo.Upsert(ctx, product)
}

func (s *basicProductService) DeleteProduct(ctx context.Context, sku string) error {
	return s.repo.Delete(ctx, sku)
}
