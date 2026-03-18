package product

import (
	"context"
	"my_project/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v5"
	"github.com/stretchr/testify/require"
)

func TestProductRepository_Create(t *testing.T) {
	ctx := context.Background()

	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewProductRepository(mock)

	name := "Apple"
	desc := "A sweet fruit"
	salePrice := 1.99
	price := 2.99
	id := uuid.New()

	rows := pgxmock.NewRows([]string{"id", "name", "description", "salePrice", "price"}).
		AddRow(id, name, desc, salePrice, price)

	mock.ExpectQuery(`INSERT INTO public\.product`).
		WithArgs(&name, &desc, &salePrice, &price).
		WillReturnRows(rows)

	got, err := repo.Create(ctx, domain.Product{
		Name:        &name,
		Description: &desc,
		SalePrice:   &salePrice,
		Price:       &price,
	})
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, id, got.Id)
	require.Equal(t, name, *got.Name)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_Update(t *testing.T) {
	ctx := context.Background()

	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewProductRepository(mock)

	id := uuid.New()
	name := "Apple Edited"
	desc := "Edited desc"
	salePrice := 3.99
	price := 4.99

	rows := pgxmock.NewRows([]string{"id", "name", "description", "salePrice", "price"}).
		AddRow(id, name, desc, salePrice, price)

	mock.ExpectQuery(`UPDATE public\.product`).
		WithArgs(id, &name, &desc, &salePrice, &price).
		WillReturnRows(rows)

	got, err := repo.Update(ctx, domain.Product{
		Id:          id,
		Name:        &name,
		Description: &desc,
		SalePrice:   &salePrice,
		Price:       &price,
	})
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, id, got.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_Get(t *testing.T) {
	ctx := context.Background()

	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewProductRepository(mock)

	id := uuid.New()
	name := "Apple"
	desc := "A sweet fruit"
	salePrice := 1.99
	price := 2.99

	rows := pgxmock.NewRows([]string{"id", "name", "description", "salePrice", "price"}).
		AddRow(id, name, desc, salePrice, price)

	mock.ExpectQuery(`SELECT id, name, description, sale_price as "salePrice", price`).
		WithArgs(id).
		WillReturnRows(rows)

	got, err := repo.Get(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, id, got.Id)

	require.NoError(t, mock.ExpectationsWereMet())
}
