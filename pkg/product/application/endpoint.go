package application

import (
	"context"
	"fmt"
	"net/url"
	"strings"

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
	Err     error                  `json:"-"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (r UpsertProductRequest) Validate() map[string]interface{} {
	errs := make(map[string]interface{})
	if r.Product == nil {

		errs["product"] = "product is required"
		return errs
	}

	r.Product.Sku = strings.TrimSpace(r.Product.Sku)
	r.Product.Name = strings.TrimSpace(r.Product.Name)
	r.Product.Brand = strings.TrimSpace(r.Product.Brand)
	r.Product.PrincipalImage = strings.TrimSpace(r.Product.PrincipalImage)

	//Validate sku
	if r.Product.Sku == "" {
		errs["sku"] = "sku is required"
	}
	if len(strings.Split(r.Product.Sku, "-")) != 2 || len(strings.Split(r.Product.Sku, "-")[1]) < 7 || len(strings.Split(r.Product.Sku, "-")[1]) > 8 {
		errs["sku"] = "sku must be in format: AAA-1234567"
	}

	//Validate name
	if r.Product.Name == "" {
		errs["name"] = "name is required"
	}
	if len(r.Product.Name) < 3 || len(r.Product.Name) > 50 {
		errs["name"] = "name must be between 3 and 50 characters"
	}

	//Validate brand
	if r.Product.Brand == "" {
		errs["brand"] = "brand is required"
	}
	if len(r.Product.Brand) < 3 || len(r.Product.Brand) > 50 {
		errs["brand"] = "brand must be between 3 and 50 characters"
	}

	//Validate size
	if r.Product.Size != nil && *r.Product.Size == "" {
		errs["size"] = "size cannot be empty"
	}

	//Validate price
	if r.Product.Price <= 1.00 || r.Product.Price >= 99999999.00 {
		errs["price"] = "price must be between 1.00 and 99999999.00"
	}

	//Validate principal_image
	if r.Product.PrincipalImage == "" {
		errs["principalImage"] = "principalImage is required"
	}

	if _, err := url.ParseRequestURI(r.Product.PrincipalImage); err != nil {
		errs["principalImage"] = "principalImage is not a valid url"
	}

	if r.Product.OtherImages != nil && len(r.Product.OtherImages) > 0 {
		for index, image := range r.Product.OtherImages {
			if _, err := url.ParseRequestURI(image); err != nil {
				errs["otherImages["+fmt.Sprint(index)+"]"] = " is not a valid url"
			}
		}
	}

	return errs
}

func (r UpsertProductResponse) Failed() error { return r.Err }

func makeUpsertProductEndpoint(svc domain.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpsertProductRequest)
		errs := req.Validate()
		if len(errs) > 0 {
			return UpsertProductResponse{Details: errs, Err: domain.ErrInvalidProduct}, nil
		}

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
