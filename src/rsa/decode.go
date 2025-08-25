package rsa

import "math/big"

// Decode decodes value using RSA algorithm. Returns an original value.
func (r *Client) Decode(value []byte) ([]byte, error) {

	if len(value) >= int(r.version) {
		return []byte{}, ErrRSAValueTooLong
	}

	var (
		origin  *big.Int = big.NewInt(0).SetBytes(value)
		decoded *big.Int = big.NewInt(0)
	)

	decoded.Exp(origin, r.keys.Private.D, r.keys.Public.N)

	return decoded.Bytes(), nil
}
