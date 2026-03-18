package usecase

import (
	"context"
	"database/sql"
	"errors"
	"my_project/internal/domain"
	orgerror "my_project/pkg/error"
	"testing"

	"github.com/google/uuid"
)

type mockProductRepo struct {
	createFn func(ctx context.Context, payload domain.Product) (*domain.Product, error)
	updateFn func(ctx context.Context, payload domain.Product) (*domain.Product, error)
	getFn    func(ctx context.Context, id uuid.UUID) (*domain.Product, error)
}

func (m *mockProductRepo) Create(ctx context.Context, payload domain.Product) (*domain.Product, error) {
	return m.createFn(ctx, payload)
}

func (m *mockProductRepo) Update(ctx context.Context, payload domain.Product) (*domain.Product, error) {
	return m.updateFn(ctx, payload)
}

func (m *mockProductRepo) Get(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	return m.getFn(ctx, id)
}

func TestProductUseCase_Create(t *testing.T) {
	t.Parallel()

	name := "Apple"
	shortName := "Ap"
	description := "sweet fruit"
	price := 100.0
	salePrice := 80.0

	tests := []struct {
		name           string
		payload        domain.Product
		mockRepo       func() domain.ProductRepository
		wantErr        bool
		wantErrCode    orgerror.Code
		wantRepoCalled bool
	}{
		{
			name: "success",
			payload: domain.Product{
				Name:        &name,
				Description: &description,
				Price:       &price,
				SalePrice:   &salePrice,
			},
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					createFn: func(ctx context.Context, payload domain.Product) (*domain.Product, error) {
						return &payload, nil
					},
				}
			},
			wantErr:        false,
			wantRepoCalled: true,
		},
		{
			name: "invalid payload",
			payload: domain.Product{
				Name:      &shortName,
				Price:     &price,
				SalePrice: &salePrice,
			},
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					createFn: func(ctx context.Context, payload domain.Product) (*domain.Product, error) {
						t.Fatalf("repo should not be called")
						return nil, nil
					},
				}
			},
			wantErr:        true,
			wantErrCode:    orgerror.CodeInvalidInput,
			wantRepoCalled: false,
		},
		{
			name: "repo error",
			payload: domain.Product{
				Name:        &name,
				Description: &description,
				Price:       &price,
				SalePrice:   &salePrice,
			},
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					createFn: func(ctx context.Context, payload domain.Product) (*domain.Product, error) {
						return nil, errors.New("db down")
					},
				}
			},
			wantErr:        true,
			wantErrCode:    orgerror.CodeSystem,
			wantRepoCalled: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewProductUseCase(tt.mockRepo())
			got, err := uc.Create(context.Background(), tt.payload)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}

				e, ok := err.(*orgerror.Error)
				if !ok {
					t.Fatalf("expected *orgerror.Error, got %T", err)
				}

				if e.Code != tt.wantErrCode {
					t.Fatalf("expected code %q, got %q", tt.wantErrCode, e.Code)
				}

				if got != nil {
					t.Fatalf("expected nil result, got %+v", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if got == nil {
				t.Fatalf("expected result, got nil")
			}
		})
	}
}

func TestProductUseCase_Update(t *testing.T) {
	t.Parallel()

	name := "Apple"
	shortName := "Ap"
	description := "sweet fruit"
	price := 100.0
	salePrice := 80.0

	tests := []struct {
		name        string
		payload     domain.Product
		mockRepo    func() domain.ProductRepository
		wantErr     bool
		wantErrCode orgerror.Code
	}{
		{
			name: "success",
			payload: domain.Product{
				Id:          uuid.New(),
				Name:        &name,
				Description: &description,
				Price:       &price,
				SalePrice:   &salePrice,
			},
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					updateFn: func(ctx context.Context, payload domain.Product) (*domain.Product, error) {
						return &payload, nil
					},
				}
			},
			wantErr: false,
		},
		{
			name: "invalid payload",
			payload: domain.Product{
				Id:        uuid.New(),
				Name:      &shortName,
				Price:     &price,
				SalePrice: &salePrice,
			},
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					updateFn: func(ctx context.Context, payload domain.Product) (*domain.Product, error) {
						t.Fatalf("repo should not be called")
						return nil, nil
					},
				}
			},
			wantErr:     true,
			wantErrCode: orgerror.CodeInvalidInput,
		},
		{
			name: "repo error",
			payload: domain.Product{
				Id:          uuid.New(),
				Name:        &name,
				Description: &description,
				Price:       &price,
				SalePrice:   &salePrice,
			},
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					updateFn: func(ctx context.Context, payload domain.Product) (*domain.Product, error) {
						return nil, errors.New("db error")
					},
				}
			},
			wantErr:     true,
			wantErrCode: orgerror.CodeSystem,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewProductUseCase(tt.mockRepo())
			got, err := uc.Update(context.Background(), tt.payload)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				e, ok := err.(*orgerror.Error)
				if !ok {
					t.Fatalf("expected *orgerror.Error, got %T", err)
				}
				if e.Code != tt.wantErrCode {
					t.Fatalf("expected code %q, got %q", tt.wantErrCode, e.Code)
				}
				if got != nil {
					t.Fatalf("expected nil result, got %+v", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if got == nil {
				t.Fatalf("expected result, got nil")
			}
		})
	}
}

func TestProductUseCase_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		id       uuid.UUID
		mockRepo func() domain.ProductRepository
		wantErr  bool
		wantCode orgerror.Code
	}{
		{
			name: "invalid id",
			id:   uuid.Nil,
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					getFn: func(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
						t.Fatalf("repo should not be called")
						return nil, nil
					},
				}
			},
			wantErr:  true,
			wantCode: orgerror.CodeInvalidInput,
		},
		{
			name: "not found",
			id:   uuid.New(),
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					getFn: func(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
						return nil, sql.ErrNoRows
					},
				}
			},
			wantErr:  true,
			wantCode: orgerror.CodeNotFound,
		},
		{
			name: "repo error",
			id:   uuid.New(),
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					getFn: func(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
						return nil, errors.New("db error")
					},
				}
			},
			wantErr:  true,
			wantCode: orgerror.CodeSystem,
		},
		{
			name: "success",
			id:   uuid.New(),
			mockRepo: func() domain.ProductRepository {
				return &mockProductRepo{
					getFn: func(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
						name := "Apple"
						price := 100.0
						salePrice := 80.0
						return &domain.Product{
							Id:        id,
							Name:      &name,
							Price:     &price,
							SalePrice: &salePrice,
						}, nil
					},
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewProductUseCase(tt.mockRepo())
			got, err := uc.Get(context.Background(), tt.id)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				e, ok := err.(*orgerror.Error)
				if !ok {
					t.Fatalf("expected *orgerror.Error, got %T", err)
				}
				if e.Code != tt.wantCode {
					t.Fatalf("expected code %q, got %q", tt.wantCode, e.Code)
				}
				if got != nil {
					t.Fatalf("expected nil result, got %+v", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if got == nil {
				t.Fatalf("expected result, got nil")
			}
		})
	}
}
