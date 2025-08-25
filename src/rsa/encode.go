package rsa

import (
	"math/big"
)

// Encode encodes value using RSA algorithm. Returns an encrypted value.
func (r *Client) Encode(value []byte) ([]byte, error) {

	if len(value) >= int(r.version)/8 {
		return []byte{}, ErrRSAValueTooLong
	}

	var (
		origin  *big.Int = big.NewInt(0).SetBytes(value)
		encoded *big.Int = big.NewInt(0)
	)

	encoded.Exp(origin, r.keys.Public.E, r.keys.Public.N)

	return encoded.Bytes(), nil
}
