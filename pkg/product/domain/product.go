package domain

type Product struct {
	Sku            string   `json:"sku"`
	Name           string   `json:"name"`
	Brand          string   `json:"brand"`
	Size           string   `json:"size"`
	Price          float64  `json:"price"`
	PrincipalImage string   `json:"principal_image"`
	OtherImages    []string `json:"other_images"`
}
