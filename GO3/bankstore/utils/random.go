package utils

import (
	"math/rand"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomAmount() pgtype.Numeric {
	return FloatToPgNumeric(float64(RandomInt(1000, 100000)) / 100)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "RUB"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
