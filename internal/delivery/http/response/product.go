package response

import "github.com/google/uuid"

type ProductResponse struct {
	Id          uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426655440000"`
	Name        string    `json:"name" example:"Apple"`
	Description *string   `json:"description" example:"A sweet fruit"`
	SalePrice   float64   `json:"salePrice" example:"1.99"`
	Price       float64   `json:"price" example:"2.99"`
}
