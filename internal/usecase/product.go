package usecase

import (
	"context"
	"database/sql"
	"errors"
	"my_project/internal/domain"
	orgerror "my_project/pkg/error"

	"github.com/google/uuid"
)

type ProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepo: productRepo}
}

func (u *ProductUseCase) Create(ctx context.Context, payload domain.Product) (*domain.Product, error) {

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	result, err := u.productRepo.Create(ctx, payload)
	if err != nil {
		return nil, orgerror.Wrap(orgerror.CodeSystem, "failed to create product", err)
	}

	return result, err
}

func (u *ProductUseCase) Update(ctx context.Context, payload domain.Product) (*domain.Product, error) {

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	result, err := u.productRepo.Update(ctx, payload)
	if err != nil {
		return nil, orgerror.Wrap(orgerror.CodeSystem, "failed to update product", err)
	}
	return result, err
}

func (u *ProductUseCase) Get(ctx context.Context, id uuid.UUID) (*domain.Product, error) {

	if id == uuid.Nil {
		return nil, orgerror.New(orgerror.CodeInvalidInput, "product id cannot be empty")
	}

	result, err := u.productRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, orgerror.New(orgerror.CodeNotFound, "product not found")
		}
		return nil, orgerror.Wrap(orgerror.CodeSystem, "failed to get product", err)
	}

	return result, err
}
