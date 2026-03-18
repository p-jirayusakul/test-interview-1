package product

import "github.com/jackc/pgx/v5/pgtype"

type productRow struct {
	Id          pgtype.UUID
	Name        pgtype.Text
	Description pgtype.Text
	SalePrice   pgtype.Numeric
	Price       pgtype.Numeric
}
