package transport_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/log"
	"github.com/jamsxd/product-crud-example/pkg/product/application"
	"github.com/jamsxd/product-crud-example/pkg/product/domain"
	"github.com/jamsxd/product-crud-example/pkg/product/infrastructure/transport"
	"github.com/jamsxd/product-crud-example/pkg/product/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHttpTransport(t *testing.T) {

	mockEndpoint := new(mocks.MockEndpoint)
	endpoint := application.ProductEndpoint{
		GetAllProducts: mockEndpoint.GetAllProducts,
		GetProduct:     mockEndpoint.GetProduct,
		UpsertProduct:  mockEndpoint.UpsertProduct,
		DeleteProduct:  mockEndpoint.DeleteProduct,
	}
	api := transport.NewHttpHandler(endpoint, log.NewJSONLogger(os.Stdout))
	srv := httptest.NewServer(api)

	t.Run("GetAllProducts", func(t *testing.T) {
		t.Run("Should return response ok", func(t *testing.T) {
			mockEndpoint.On("GetAllProducts", mock.Anything, application.GetAllProductsRequest{}).Return(&application.GetAllProductsResponse{}, nil).Once()
			req, _ := http.NewRequest("GET", srv.URL+"/products", nil)
			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, http.StatusOK, res.StatusCode)
		})
	})

	t.Run("GetProduct", func(t *testing.T) {
		t.Run("Should return response ok", func(t *testing.T) {
			mockEndpoint.On("GetProduct", mock.Anything, application.GetProductRequest{Sku: "FAL-123123"}).Return(&application.GetProductResponse{}, nil).Once()
			req, _ := http.NewRequest("GET", srv.URL+"/products/FAL-123123", nil)
			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, http.StatusOK, res.StatusCode)
		})
		t.Run("Should return not found", func(t *testing.T) {
			mockEndpoint.On("GetProduct", mock.Anything, mock.Anything).Return(&application.GetProductResponse{Err: domain.ErrProductNotFound}, nil).Once()
			req, _ := http.NewRequest("GET", srv.URL+"/products/FAL-123123", nil)
			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, http.StatusNotFound, res.StatusCode)
		})
	})

	t.Run("UpsertProduct", func(t *testing.T) {
		t.Run("Should return response ok", func(t *testing.T) {
			mockEndpoint.On("UpsertProduct", mock.Anything, application.UpsertProductRequest{
				Product: &domain.Product{
					Sku:   "FAL-123123",
					Name:  "FAL-123123",
					Brand: "FAL-123123",
					Price: 123.123,
				},
			}).Return(&application.UpsertProductResponse{}, nil).Once()
			body := `{"sku":"FAL-123123","name":"FAL-123123","brand":"FAL-123123","price":123.123}`
			req, _ := http.NewRequest("PUT", srv.URL+"/products", strings.NewReader(body))
			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, http.StatusOK, res.StatusCode)
		})
		t.Run("Should return bad request response", func(t *testing.T) {
			mockEndpoint.On("UpsertProduct", mock.Anything, application.UpsertProductRequest{
				Product: &domain.Product{
					Sku:   "FAL-123123",
					Name:  "FAL-123123",
					Brand: "FAL-123123",
					Price: 123.123,
				},
			}).Return(&application.UpsertProductResponse{Err: domain.ErrInvalidProduct}, nil).Once()
			body := `{"sku":"FAL-123123","name":"FAL-123123","brand":"FAL-123123","price":123.123}`
			req, _ := http.NewRequest("PUT", srv.URL+"/products", strings.NewReader(body))
			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		})
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		t.Run("Should return response ok", func(t *testing.T) {
			mockEndpoint.On("DeleteProduct", mock.Anything, application.DeleteProductRequest{Sku: "FAL-123123"}).Return(&application.DeleteProductResponse{}, nil).Once()
			req, _ := http.NewRequest("DELETE", srv.URL+"/products/FAL-123123", nil)
			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, http.StatusOK, res.StatusCode)
		})
	})

}
