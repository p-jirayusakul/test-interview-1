package product

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type productRow struct {
	Id          uuid.UUID
	Name        string
	Description pgtype.Text
	SalePrice   float64
	Price       float64
}
