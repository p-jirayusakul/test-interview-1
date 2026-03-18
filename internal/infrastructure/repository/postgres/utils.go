package postgres

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func PgUUIDToUUID(v pgtype.UUID) uuid.UUID {
	if !v.Valid {
		return uuid.Nil
	}

	id, err := uuid.Parse(v.String())
	if err != nil {
		return uuid.Nil
	}
	return id
}

func PgTextToPtString(v pgtype.Text) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

func PgNumToFloat64(v pgtype.Numeric) *float64 {
	if v.Valid {
		vf, err := v.Float64Value()
		if err != nil {
			return nil
		}
		return &vf.Float64
	}

	return nil
}
