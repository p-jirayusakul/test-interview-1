package product

import (
	"my_project/internal/domain"

	pg "my_project/internal/infrastructure/repository/postgres"
)

func mappingRowToDomain(v *productRow) *domain.Product {
	return &domain.Product{
		Id:          v.Id,
		Name:        &v.Name,
		Description: pg.PgTextToPtString(v.Description),
		SalePrice:   &v.SalePrice,
		Price:       &v.Price,
	}
}
