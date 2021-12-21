package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jamsxd/product-crud-example/pkg/product/application"
	"github.com/jamsxd/product-crud-example/pkg/product/domain"
)

func NewHttpHandler(endpoints application.ProductEndpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/products").Handler(
		kithttp.NewServer(
			endpoints.GetAllProducts,
			decodeGetAllProductsRequest,
			encodeHTTPGenericResponse,
			kithttp.ServerErrorLogger(logger),
		),
	)
	r.Methods("GET").Path("/products/{sku}").Handler(
		kithttp.NewServer(
			endpoints.GetProduct,
			decodeGetProductRequest,
			encodeHTTPGenericResponse,
			kithttp.ServerErrorLogger(logger),
		),
	)
	r.Methods("PUT").Path("/products").Handler(
		kithttp.NewServer(
			endpoints.UpsertProduct,
			decodeUpsertProductRequest,
			encodeUpsertResponse,
			kithttp.ServerErrorLogger(logger),
		),
	)
	r.Methods("DELETE").Path("/products/{sku}").Handler(
		kithttp.NewServer(
			endpoints.DeleteProduct,
			decodeDeleteProductRequest,
			encodeHTTPGenericResponse,
			kithttp.ServerErrorLogger(logger),
		),
	)
	return handlers.RecoveryHandler()(r)
}

func decodeGetAllProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return application.GetAllProductsRequest{}, nil
}

func decodeGetProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	sku := mux.Vars(r)["sku"]
	return application.GetProductRequest{Sku: sku}, nil
}

func decodeUpsertProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return nil, err
	}
	return application.UpsertProductRequest{Product: &product}, nil
}

func encodeUpsertResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		w.WriteHeader(err2code(f.Failed()))
		json.NewEncoder(w).Encode(response)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	sku := mux.Vars(r)["sku"]
	return application.DeleteProductRequest{Sku: sku}, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	switch err {
	case domain.ErrInvalidProduct:
		return http.StatusBadRequest
	case domain.ErrProductNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

type errorWrapper struct {
	Error string `json:"error"`
}
