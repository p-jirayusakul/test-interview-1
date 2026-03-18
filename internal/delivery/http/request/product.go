package request

type UpdateProductRequest struct {
	Name        *string  `json:"name" binding:"omitempty,min=3" example:"Apple-Edited"`
	Description *string  `json:"description" example:"A sweet fruit edited"`
	SalePrice   *float64 `json:"salePrice" binding:"omitempty,gt=0" example:"1.99"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0" example:"2.99"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=3" example:"Apple"`
	Description *string `json:"description" example:"A sweet fruit"`
	SalePrice   float64 `json:"salePrice" binding:"required,gt=0" example:"1.99"`
	Price       float64 `json:"price" binding:"required,gt=0" example:"2.99"`
}
