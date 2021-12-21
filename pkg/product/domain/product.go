package domain

type Product struct {
	Sku            string   `json:"sku"`
	Name           string   `json:"name"`
	Brand          string   `json:"brand"`
	Size           *string  `json:"size"`
	Price          float64  `json:"price"`
	PrincipalImage string   `json:"principalImage"`
	OtherImages    []string `json:"otherImages"`
}
