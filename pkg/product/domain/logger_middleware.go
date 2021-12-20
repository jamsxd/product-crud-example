package domain

import (
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

func (mw LoggingMiddleware) GetProduct(sku string) (*Product, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetProduct",
			"sku", sku,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetProduct(sku)
}

func (mw LoggingMiddleware) UpsertProduct(sku string, product *Product) error {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Upsert",
			"sku", sku,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.UpsertProduct(sku, product)
}

func (mw LoggingMiddleware) DeleteProduct(sku string) error {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Delete",
			"sku", sku,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.DeleteProduct(sku)
}
