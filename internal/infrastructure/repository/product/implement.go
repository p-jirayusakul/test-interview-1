package product

import (
	"context"
	"my_project/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) domain.ProductRepository {
	return &productRepo{pool: pool}
}

func (r *productRepo) Create(ctx context.Context, payload domain.Product) (*domain.Product, error) {

	var row productRow

	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO public.product (name, description, sale_price, price)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, name, description, sale_price as "salePrice", price`,
		payload.Name, payload.Description, payload.SalePrice, payload.Price,
	).Scan(&row.Id, &row.Name, &row.Description, &row.SalePrice, &row.Price)

	if err != nil {
		return nil, err
	}

	return mappingRowToDomain(&row), nil
}

func (r *productRepo) Update(ctx context.Context, payload domain.Product) (*domain.Product, error) {

	var row productRow

	err := r.pool.QueryRow(
		ctx,
		`UPDATE public.product
		 SET name = COALESCE($2, name),
		     description = COALESCE($3, description),
		     sale_price = COALESCE($4, sale_price),
		     price = COALESCE($5, price),
		     updated_at = now()
		 WHERE id = $1 RETURNING id, name, description, sale_price as "salePrice", price`,
		payload.Id, payload.Name, payload.Description, payload.SalePrice, payload.Price,
	).Scan(&row.Id, &row.Name, &row.Description, &row.SalePrice, &row.Price)

	if err != nil {
		return nil, err
	}

	return mappingRowToDomain(&row), nil
}

func (r *productRepo) Get(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var row productRow

	err := r.pool.QueryRow(
		ctx,
		`SELECT id, name, description, sale_price as "salePrice", price  FROM public.product WHERE id=$1`,
		id,
	).Scan(&row.Id, &row.Name, &row.Description, &row.SalePrice, &row.Price)

	if err != nil {
		return nil, err
	}

	return mappingRowToDomain(&row), nil
}
