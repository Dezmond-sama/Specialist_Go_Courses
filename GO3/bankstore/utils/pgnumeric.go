package utils

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

func FloatToPgNumeric(f float64) pgtype.Numeric {
	return pgtype.Numeric{
		NaN:   false,
		Valid: true,
		Int:   big.NewInt(int64(f * 100)),
		Exp:   -2,
	}
}

func PgNumericToFloat(n pgtype.Numeric) float64 {
	return float64(n.Int.Int64()) / 100
}
