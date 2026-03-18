package product

import (
	"my_project/internal/domain"

	pg "my_project/internal/infrastructure/repository/postgres"
)

func mappingRowToDomain(v *productRow) *domain.Product {
	return &domain.Product{
		Id:          pg.PgUUIDToUUID(v.Id),
		Name:        pg.PgTextToPtString(v.Name),
		Description: pg.PgTextToPtString(v.Description),
		SalePrice:   pg.PgNumToFloat64(v.SalePrice),
		Price:       pg.PgNumToFloat64(v.Price),
	}
}
