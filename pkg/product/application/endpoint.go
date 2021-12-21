package application

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jamsxd/product-crud-example/pkg/product/domain"
)

type ProductEndpoint struct {
	GetAllProducts endpoint.Endpoint
	GetProduct     endpoint.Endpoint
	UpsertProduct  endpoint.Endpoint
	DeleteProduct  endpoint.Endpoint
}

func NewProductEndpoint(svc domain.ProductService) ProductEndpoint {
	return ProductEndpoint{
		GetAllProducts: makeGetAllProductsEndpoint(svc),
		GetProduct:     makeGetProductEndpoint(svc),
		UpsertProduct:  makeUpsertProductEndpoint(svc),
		DeleteProduct:  makeDeleteProductEndpoint(svc),
	}
}

type GetAllProductsRequest struct{}

type GetAllProductsResponse struct {
	Products []domain.Product `json:"products"`
	Err      error            `json:"err,omitempty"`
}

func (r GetAllProductsResponse) Failed() error { return r.Err }

func makeGetAllProductsEndpoint(svc domain.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := svc.GetAllProducts(ctx)
		return GetAllProductsResponse{Products: products, Err: err}, nil
	}
}

type GetProductRequest struct {
	Sku string `json:"sku"`
}

type GetProductResponse struct {
	Product *domain.Product `json:"product"`
	Err     error           `json:"err,omitempty"`
}

func (r GetProductResponse) Failed() error { return r.Err }

func makeGetProductEndpoint(svc domain.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		product, err := svc.GetProduct(ctx, req.Sku)
		return GetProductResponse{Product: product, Err: err}, nil
	}
}

type UpsertProductRequest struct {
	Product *domain.Product `json:"product"`
}

type UpsertProductResponse struct {
	Err error `json:"err,omitempty"`
}

func (r UpsertProductResponse) Failed() error { return r.Err }

func makeUpsertProductEndpoint(svc domain.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpsertProductRequest)
		err := svc.UpsertProduct(ctx, req.Product)
		return UpsertProductResponse{Err: err}, nil
	}
}

type DeleteProductRequest struct {
	Sku string `json:"sku"`
}

type DeleteProductResponse struct {
	Err error `json:"err,omitempty"`
}

func (r DeleteProductResponse) Failed() error { return r.Err }

func makeDeleteProductEndpoint(svc domain.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		err := svc.DeleteProduct(ctx, req.Sku)
		return DeleteProductResponse{Err: err}, nil
	}
}
