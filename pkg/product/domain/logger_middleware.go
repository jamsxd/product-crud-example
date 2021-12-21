package domain

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type Middleware func(ProductService) ProductService

type LoggingMiddleware struct {
	logger log.Logger
	next   ProductService
}

func newLoggingMiddleware(logger log.Logger) Middleware {
	return func(next ProductService) ProductService {
		return LoggingMiddleware{logger, next}
	}
}

func (mw LoggingMiddleware) GetAllProducts(ctx context.Context) ([]Product, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetAllProducts",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetAllProducts(ctx)
}

func (mw LoggingMiddleware) GetProduct(ctx context.Context, sku string) (*Product, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetProduct",
			"sku", sku,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetProduct(ctx, sku)
}

func (mw LoggingMiddleware) UpsertProduct(ctx context.Context, product *Product) error {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Upsert",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.UpsertProduct(ctx, product)
}

func (mw LoggingMiddleware) DeleteProduct(ctx context.Context, sku string) error {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Delete",
			"sku", sku,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.DeleteProduct(ctx, sku)
}
