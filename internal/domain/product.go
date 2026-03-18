package domain

import (
	"context"
	orgerror "my_project/pkg/error"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID
	Name        *string
	Description *string
	SalePrice   *float64
	Price       *float64
}

func (p *Product) Validate() error {
	if p.Name != nil && len(*p.Name) < 3 {
		return orgerror.New(orgerror.CodeInvalidInput, "product name must be at least 3 characters")
	}

	if p.Price != nil && *p.Price <= 0 {
		return orgerror.New(orgerror.CodeInvalidInput, "product price must be greater than 0")
	}

	if p.SalePrice != nil && *p.SalePrice <= 0 {
		return orgerror.New(orgerror.CodeInvalidInput, "product sale price must be greater than 0")
	}

	return nil
}

type ProductRepository interface {
	Create(ctx context.Context, payload Product) (*Product, error)
	Update(ctx context.Context, product Product) (*Product, error)
	Get(ctx context.Context, id uuid.UUID) (*Product, error)
}
