package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func PgTextToPtString(v pgtype.Text) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}
