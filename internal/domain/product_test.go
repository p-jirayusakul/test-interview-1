package domain

import (
	"errors"
	"testing"

	orgerror "my_project/pkg/error"

	"github.com/google/uuid"
)

func TestProduct_Validate(t *testing.T) {
	t.Parallel()

	name := "Apple"
	shortName := "Ap"
	price := 10.0
	zeroPrice := 0.0
	negativePrice := -1.0
	salePrice := 5.0
	zeroSalePrice := 0.0
	negativeSalePrice := -1.0

	tests := []struct {
		name     string
		product  Product
		wantErr  bool
		wantCode orgerror.Code
		wantMsg  string
	}{
		{
			name: "valid product",
			product: Product{
				Id:          uuid.New(),
				Name:        &name,
				Price:       &price,
				SalePrice:   &salePrice,
				Description: nil,
			},
			wantErr: false,
		},
		{
			name: "invalid name length",
			product: Product{
				Name:      &shortName,
				Price:     &price,
				SalePrice: &salePrice,
			},
			wantErr:  true,
			wantCode: orgerror.CodeInvalidInput,
			wantMsg:  "product name must be at least 3 characters",
		},
		{
			name: "invalid price zero",
			product: Product{
				Name:      &name,
				Price:     &zeroPrice,
				SalePrice: &salePrice,
			},
			wantErr:  true,
			wantCode: orgerror.CodeInvalidInput,
			wantMsg:  "product price must be greater than 0",
		},
		{
			name: "invalid price negative",
			product: Product{
				Name:      &name,
				Price:     &negativePrice,
				SalePrice: &salePrice,
			},
			wantErr:  true,
			wantCode: orgerror.CodeInvalidInput,
			wantMsg:  "product price must be greater than 0",
		},
		{
			name: "invalid sale price zero",
			product: Product{
				Name:      &name,
				Price:     &price,
				SalePrice: &zeroSalePrice,
			},
			wantErr:  true,
			wantCode: orgerror.CodeInvalidInput,
			wantMsg:  "product sale price must be greater than 0",
		},
		{
			name: "invalid sale price negative",
			product: Product{
				Name:      &name,
				Price:     &price,
				SalePrice: &negativeSalePrice,
			},
			wantErr:  true,
			wantCode: orgerror.CodeInvalidInput,
			wantMsg:  "product sale price must be greater than 0",
		},
		{
			name: "nil fields are allowed by validate",
			product: Product{
				Id: uuid.New(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.product.Validate()

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}

				var gotErr *orgerror.Error
				ok := errors.As(err, &gotErr)
				if !ok {
					t.Fatalf("expected *orgerror.Error, got %T", err)
				}

				if gotErr.Code != tt.wantCode {
					t.Fatalf("expected code %q, got %q", tt.wantCode, gotErr.Code)
				}

				if gotErr.Message != tt.wantMsg {
					t.Fatalf("expected message %q, got %q", tt.wantMsg, gotErr.Message)
				}

				return
			}

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
		})
	}
}
