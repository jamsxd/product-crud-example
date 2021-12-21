package mocks

import (
	"context"

	"github.com/jamsxd/product-crud-example/pkg/product/domain"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *MockService) GetProduct(ctx context.Context, sku string) (*domain.Product, error) {
	args := m.Called(ctx, sku)
	var r0 *domain.Product
	if rf, ok := args.Get(0).(func(context.Context, string) *domain.Product); ok {
		r0 = rf(ctx, sku)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sku)
	} else {
		r1 = args.Error(1)
	}
	return r0, r1
}

func (m *MockService) UpsertProduct(ctx context.Context, product *domain.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockService) DeleteProduct(ctx context.Context, sku string) error {
	args := m.Called(ctx, sku)
	return args.Error(0)
}
