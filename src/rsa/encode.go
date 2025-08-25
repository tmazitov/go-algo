package rsa

import (
	"math/big"
)

// Encode encodes value using RSA algorithm. Returns an encrypted value.
func (r *RSA) Encode(value string) *big.Int {
	var (
		valueBytes []byte   = []byte(value)
		origin     *big.Int = big.NewInt(0).SetBytes(valueBytes)
		encoded    *big.Int = big.NewInt(0)
	)

	encoded.Exp(origin, r.public.E, r.public.N)

	return encoded
}
