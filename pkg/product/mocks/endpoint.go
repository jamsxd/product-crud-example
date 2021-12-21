package mocks

import (
	"context"

	"github.com/jamsxd/product-crud-example/pkg/product/application"
	"github.com/stretchr/testify/mock"
)

type MockEndpoint struct {
	mock.Mock
}

func (m *MockEndpoint) GetAllProducts(ctx context.Context, request interface{}) (interface{}, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*application.GetAllProductsResponse), args.Error(1)
}

func (m *MockEndpoint) GetProduct(ctx context.Context, request interface{}) (interface{}, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*application.GetProductResponse), args.Error(1)
}

func (m *MockEndpoint) UpsertProduct(ctx context.Context, request interface{}) (interface{}, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*application.UpsertProductResponse), args.Error(1)
}

func (m *MockEndpoint) DeleteProduct(ctx context.Context, request interface{}) (interface{}, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*application.DeleteProductResponse), args.Error(1)
}
