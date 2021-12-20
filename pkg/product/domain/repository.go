package domain

type ProductRepository interface {
	FindBySku(sku string) (*Product, error)
	Upsert(sku string, product *Product) error
	Delete(sku string) error
}
