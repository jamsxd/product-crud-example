package domain

import "github.com/go-kit/kit/log"

type ProductService interface {
	GetProduct(sku string) (*Product, error)
	UpsertProduct(sku string, product *Product) error
	DeleteProduct(sku string) error
}

type basicProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository, logger log.Logger) ProductService {
	return newLoggingMiddleware(logger)(&basicProductService{repo})
}

func (s *basicProductService) GetProduct(sku string) (*Product, error) {
	return s.repo.FindBySku(sku)
}

func (s *basicProductService) UpsertProduct(sku string, product *Product) error {
	return s.repo.Upsert(sku, product)
}

func (s *basicProductService) DeleteProduct(sku string) error {
	return s.repo.Delete(sku)
}
