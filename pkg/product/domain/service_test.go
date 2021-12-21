package domain_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"
	"github.com/jamsxd/product-crud-example/pkg/product/domain"
	"github.com/jamsxd/product-crud-example/pkg/product/mocks"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {

	repository := new(mocks.MockRepository)
	svc := domain.NewProductService(repository, log.NewJSONLogger(os.Stdout))

	t.Run("GetAllProducts", func(t *testing.T) {
		t.Run("Should return all products", func(t *testing.T) {
			// Given
			repository.On("FindAll", mock.Anything).Return([]domain.Product{}, nil)

			// When
			products, err := svc.GetAllProducts(context.Background())

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}
			if len(products) != 0 {
				t.Errorf("Expected products to be empty, got %v", products)
			}
		})
	})

	t.Run("GetProduct", func(t *testing.T) {
		t.Run("Should return product", func(t *testing.T) {
			// Given
			repository.On("FindBySku", mock.Anything, "sku").Return(
				&domain.Product{
					Sku: "sku",
				},
				nil,
			).Once()

			// When
			product, err := svc.GetProduct(context.Background(), "sku")

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}
			if product.Sku != "sku" {
				t.Errorf("Expected product to be sku, got %v", product.Sku)
			}
		})
		t.Run("Should return an error", func(t *testing.T) {
			// Given
			repository.On("FindBySku", mock.Anything, "sku").Return(nil, domain.ErrProductNotFound).Once()

			// When
			product, err := svc.GetProduct(context.Background(), "sku")

			// Then
			if err == nil {
				t.Errorf("Expected error to be not nil, got %v", err)
			}
			if product != nil {
				t.Errorf("Expected product to be nil, got %v", product)
			}
		})
	})

	t.Run("CreateProduct", func(t *testing.T) {
		t.Run("Should create product", func(t *testing.T) {
			// Given
			repository.On("Upsert", mock.Anything, &domain.Product{}).Return(nil)

			// When
			err := svc.UpsertProduct(context.Background(), &domain.Product{})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}
		})
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		t.Run("Should delete product", func(t *testing.T) {
			// Given
			repository.On("Delete", mock.Anything, "sku").Return(nil)

			// When
			err := svc.DeleteProduct(context.Background(), "sku")

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}
		})
	})
}
