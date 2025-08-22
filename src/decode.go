package src

import "math/big"

func (r *RSA) Decode(value int64) int64 {
	var (
		origin  *big.Int = big.NewInt(value)
		dencoded *big.Int = big.NewInt(0)
	)

	dencoded.Exp(origin, r.private.D, r.public.N)

	return dencoded.Int64()
}
