package rsa

import (
	"errors"
	"math/big"
)

type Client struct {
	keys    *KeyPair
	version RSAVersion
}

// NewClient creates a Client that can be used for encoding and decoding using RSA algorithm.
// Client structure don't share public and private keys. For this reason, if you want to decode your encoded message,
// you have to use the same client to get a correct original data.
func NewClient(version RSAVersion, e int64) (*Client, error) {

	var E *big.Int = big.NewInt(e)

	// Negative or zero
	if E.Cmp(big.NewInt(0)) <= 0 {
		return nil, ErrRSAInvalidArg
	}

	var (
		base *rsaBase
		err  error
	)

	if base, err = setupEnvirement(version, E); err != nil {
		return nil, errors.Join(ErrRSAIntenal, err)
	}

	var instance *Client = &Client{
		version: version,
		keys: &KeyPair{
			Public: &PublicKey{
				N: base.N,
				E: E,
			},
			Private: &PrivateKey{
				D: base.d,
			},
		},
	}

	return instance, nil
}
