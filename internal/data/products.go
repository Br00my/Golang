package data

import (
	"greenworld.wow/internal/validator"
)

type Product struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Price int8 `json:"price"`
	Category string `json:"category"`
}

func ValidateProduct(v *validator.Validator, product *Product) {
	v.Check(product.Title != "", "title", "must be provided")
	v.Check(len(product.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(product.Category == "Reusable Water Bottles" || product.Category == "Bamboo Toothbrushes" || product.Category == "Solar Charges" || product.Category == "Eco-Friendly Cleaning Products" || product.Category == "Recycled Paper Products", "category", "is invalid")
	v.Check(product.Price > 0, "price", "must be provided")
}
