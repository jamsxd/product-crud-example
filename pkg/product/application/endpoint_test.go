package application_test

import (
	"context"
	"testing"

	"github.com/jamsxd/product-crud-example/pkg/product/application"
	"github.com/jamsxd/product-crud-example/pkg/product/domain"
	"github.com/jamsxd/product-crud-example/pkg/product/mocks"
	"github.com/stretchr/testify/mock"
)

func TestEndpoints(t *testing.T) {
	svc := new(mocks.MockService)
	endpoints := application.NewProductEndpoint(svc)

	t.Run("GetAllProducts", func(t *testing.T) {
		t.Run("Should return all products", func(t *testing.T) {
			// Given
			svc.On("GetAllProducts", mock.Anything).Return([]domain.Product{}, nil)

			// When
			products, err := endpoints.GetAllProducts(context.Background(), application.GetAllProductsRequest{})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if products.(application.GetAllProductsResponse).Failed() != nil {
				t.Errorf("Expected products err to be nil, got %v", products.(application.GetAllProductsResponse).Failed())
			}

			if products.(application.GetAllProductsResponse).Products == nil {
				t.Errorf("Expected products to be not nil, got %v", products)
			}
		})
	})

	t.Run("GetProduct", func(t *testing.T) {
		t.Run("Should return product", func(t *testing.T) {
			// Given
			svc.On("GetProduct", mock.Anything, mock.Anything).Return(&domain.Product{}, nil)

			// When
			product, err := endpoints.GetProduct(context.Background(), application.GetProductRequest{Sku: "sku"})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if product.(application.GetProductResponse).Failed() != nil {
				t.Errorf("Expected product err to be nil, got %v", product.(application.GetProductResponse).Err)
			}

			if product.(application.GetProductResponse).Product == nil {
				t.Errorf("Expected product to be not nil, got %v", product)
			}
		})
	})

	t.Run("CreateProduct", func(t *testing.T) {
		t.Run("Should return bad request error", func(t *testing.T) {
			// When
			product, err := endpoints.UpsertProduct(context.Background(), application.UpsertProductRequest{})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if product.(application.UpsertProductResponse).Failed() == nil {
				t.Errorf("Expected product err to be not nil, got %v", product.(application.UpsertProductResponse).Err)
			}
			if product.(application.UpsertProductResponse).Details == nil {
				t.Errorf("Expected product details to be not nil, got %v", product.(application.UpsertProductResponse).Details)
			}

		})

		t.Run("Should return bad request error by empty fields", func(t *testing.T) {
			// When
			product, err := endpoints.UpsertProduct(context.Background(), application.UpsertProductRequest{
				Product: &domain.Product{},
			})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if product.(application.UpsertProductResponse).Failed() == nil {
				t.Errorf("Expected product err to be not nil, got %v", product.(application.UpsertProductResponse).Err)
			}
		})

		t.Run("Should return bad request error by invalid format", func(t *testing.T) {
			// When
			size := ""
			product, err := endpoints.UpsertProduct(context.Background(), application.UpsertProductRequest{
				Product: &domain.Product{
					Size:           &size,
					Sku:            "sku-123123123123",
					Name:           "na",
					Brand:          "br",
					Price:          1.0,
					PrincipalImage: "http//principalImage",
					OtherImages:    []string{"http//image1", "http//image2"},
				},
			})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if product.(application.UpsertProductResponse).Failed() == nil {
				t.Errorf("Expected product err to be not nil, got %v", product.(application.UpsertProductResponse).Err)
			}
		})

		t.Run("Should create product", func(t *testing.T) {
			// Given
			svc.On("UpsertProduct", mock.Anything, mock.Anything).Return(nil)

			// When
			size := "40"
			product, err := endpoints.UpsertProduct(context.Background(), application.UpsertProductRequest{
				Product: &domain.Product{
					Sku:            "FAL-8406270",
					Name:           "500 Zapatilla Urbana Mujer",
					Brand:          "New Balance",
					Size:           &size,
					Price:          200.00,
					PrincipalImage: "https://falabella.scene7.com/is/image/Falabella/8406270_1",
				},
			})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if product.(application.UpsertProductResponse).Failed() != nil {
				t.Errorf("Expected product err to be nil, got %v", product)
			}
		})
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		t.Run("Should delete product", func(t *testing.T) {
			// Given
			svc.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil)

			// When
			res, err := endpoints.DeleteProduct(context.Background(), application.DeleteProductRequest{Sku: "sku"})

			// Then
			if err != nil {
				t.Errorf("Expected error to be nil, got %v", err)
			}
			if res.(application.DeleteProductResponse).Failed() != nil {
				t.Errorf("Expected product err to be nil, got %v", res)
			}
		})
	})
}
