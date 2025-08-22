package src

import (
	"math/big"
)

type Encoded struct {
	Key []byte
}

func (r *RSA) Encode(value int64) int64 {
	var (
		origin  *big.Int = big.NewInt(value)
		encoded *big.Int = big.NewInt(0)
	)

	encoded.Exp(origin, r.public.E, r.public.N)

	return encoded.Int64()
}
