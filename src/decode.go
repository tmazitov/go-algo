package src

import "math/big"

// Decode decodes value using RSA algorithm. Returns an original value.
func (r *RSA) Decode(value *big.Int) string {
	var (
		decoded *big.Int = big.NewInt(0)
		decodedValue string
	)

	decoded.Exp(value, r.private.D, r.public.N)
	decodedValue = string(decoded.Bytes())

	return decodedValue
}
